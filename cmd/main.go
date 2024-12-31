package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dylanmazurek/go-lunchmoney/pkg/lunchmoney"
	"github.com/dylanmazurek/go-lunchmoney/pkg/lunchmoney/models"
	"github.com/dylanmazurek/go-lunchmoney/pkg/utilities/vault"
	"github.com/markkurossi/tabulate"
)

func main() {
	ctx := context.Background()

	vaultAddr := os.Getenv("VAULT_ADDR")
	vaultAppRoleId := os.Getenv("VAULT_APP_ROLE_ID")
	vaultSecretId := os.Getenv("VAULT_SECRET_ID")

	vaultClient, err := vault.NewClient(ctx, vaultAddr, vaultAppRoleId, vaultSecretId)
	if err != nil {
		panic(err)
	}

	flags := getFlags()
	client, err := lunchmoney.New(ctx, lunchmoney.WithVaultClient(vaultClient))
	if err != nil {
		panic(err)
	}

	start, _ := time.Parse("2006-01-02", flags.StartDate)
	end, _ := time.Parse("2006-01-02", flags.EndDate)

	st := &lunchmoney.ListTransactionFilter{
		StartDate: start,
		EndDate:   end,
	}

	transactions, err := client.ListTransaction(*st)
	if err != nil {
		panic(err)
	}

	tab := tabulate.New(tabulate.CompactUnicode)
	tab.Header("ID")
	tab.Header("Original Name")
	tab.Header("Display Name")
	tab.Header("Account Name")
	tab.Header("Amount")
	tab.Header("Category")
	tab.Header("Tags")

	for _, transaction := range *transactions {
		row := tab.Row()
		row.Column(transaction.ID.String())
		row.Column(transaction.OriginalName)
		row.Column(transaction.DisplayName)
		row.Column(transaction.AssetDisplayName)
		row.Column(fmt.Sprintf("%.2f", transaction.Amount.AsMajorUnits()))
		if transaction.CategoryName == nil {
			row.Column("")
		} else {
			row.Column(*transaction.CategoryName)
		}
		row.Column(strings.Join(transaction.Tags, ","))

		fmt.Println(tab.String())
	}

}

func getFlags() models.Flags {
	flags := models.Flags{}
	today := time.Now()
	flag.StringVar(&flags.AssetID, "assetId", "", "id of the lunchmoney asset account")
	flag.StringVar(&flags.StartDate, "startDate", today.AddDate(0, 0, -100).Format("2006-01-02"), "start date of the transaction range")
	flag.StringVar(&flags.EndDate, "endDate", "", "end date of the transaction range")

	flag.Parse()

	return flags
}
