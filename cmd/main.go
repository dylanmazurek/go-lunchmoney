package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/dylanmazurek/go-lunchmoney/pkg/lunchmoney"
	"github.com/dylanmazurek/go-lunchmoney/pkg/lunchmoney/models"
	"github.com/dylanmazurek/go-lunchmoney/pkg/utilities/truncate"
	"github.com/markkurossi/tabulate"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()
	initLogger()

	flags := getFlags()

	lunchmoneyOpts := []lunchmoney.Option{
		lunchmoney.WithAPIKey(flags.LunchmoneyAPIKey),
	}

	client, err := lunchmoney.New(ctx, lunchmoneyOpts...)
	if err != nil {
		panic(err)
	}

	st := &lunchmoney.ListTransactionFilter{
		AssetID:   &flags.AssetID,
		StartDate: flags.StartDate,
		EndDate:   flags.EndDate,
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

	for _, transaction := range *transactions {
		row := tab.Row()
		row.Column(transaction.ID.String())
		row.Column(truncate.TruncateText(transaction.OriginalName, 30))
		row.Column(transaction.DisplayName)
		row.Column(transaction.AssetDisplayName)
		row.Column(fmt.Sprintf("%.2f", transaction.Amount.AsMajorUnits()))
		if transaction.CategoryName == nil {
			row.Column("")
		} else {
			row.Column(*transaction.CategoryName)
		}
	}

	fmt.Println(tab.String())
}

func initLogger() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	log.Logger = logger
}

func getFlags() models.Flags {
	var flags struct {
		LunchmoneyAPIKey string
		AssetID          string
		StartDate        string
		EndDate          string
	}

	flag.StringVar(&flags.LunchmoneyAPIKey, "apiKey", "", "lunchmoney api key")
	flag.StringVar(&flags.AssetID, "assetId", "", "id of the lunchmoney asset account")
	flag.StringVar(&flags.StartDate, "startDate", "", "start date of the transaction range")
	flag.StringVar(&flags.EndDate, "endDate", "", "end date of the transaction range")

	flag.Parse()

	var parsedFlags models.Flags

	parsedFlags.LunchmoneyAPIKey = flags.LunchmoneyAPIKey
	if parsedFlags.LunchmoneyAPIKey == "" {
		lunchmoneyAPIKey, envExists := os.LookupEnv("LUNCHMONEY_API_KEY")
		if envExists {
			parsedFlags.LunchmoneyAPIKey = lunchmoneyAPIKey
			log.Info().Msg("using LUNCHMONEY_API_KEY environment variable")
		}
	}

	apiKeyRegex := regexp.MustCompile(`(?i)^[a-z0-9]{50}$`)
	isValidKey := apiKeyRegex.MatchString(parsedFlags.LunchmoneyAPIKey)
	if !isValidKey {
		log.Panic().Msg("apiKey must be a 50 character alphanumeric string")
	}

	assetId, err := strconv.ParseInt(flags.AssetID, 10, 64)
	if err != nil {
		log.Panic().Msg("assetId must be an integer")
	}
	parsedFlags.AssetID = assetId

	startDate, err := time.Parse("2006-01-02", flags.StartDate)
	if err != nil {
		log.Panic().Msg("startDate must be in the format yyyy-mm-dd")
	}
	parsedFlags.StartDate = startDate

	endDate, err := time.Parse("2006-01-02", flags.EndDate)
	if err != nil {
		log.Panic().Msg("endDate must be in the format yyyy-mm-dd")
	}
	parsedFlags.EndDate = endDate

	return parsedFlags
}
