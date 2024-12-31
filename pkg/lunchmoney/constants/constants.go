package constants

const (
	PLUGIN_NAME  = "lunchmoney"
	REPO_URL     = "github.com/dylanmazurek/go-lunchmoney"
	API_BASE_URL = "https://dev.lunchmoney.app/v1"
)

const (
	API_PATH_ME           = "/me"
	API_PATH_ASSETS       = "/assets"
	API_PATH_CATEGORIES   = "/categories"
	API_PATH_TRANSACTIONS = "/transactions"
	API_PATH_TAGS         = "/tags"
)

const (
	STATE_NEW           = "NEW"
	STATE_INITIALIZED   = "INITIALIZED"
	STATE_AUTHENTICATED = "AUTHENTICATED"
	STATE_ERROR         = "ERROR"
)
