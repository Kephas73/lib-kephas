package error_code

const (
	ERROR_PASSWORD_LEVEL2 int32 = 202001
	ERROR_GOOGLE_AUTH     int32 = 202002

	ERROR_CODE_OK int32 = 0

	ERROR_BAD_REQUEST    int32 = 400000
	ERROR_BIND_DATA      int32 = 400001
	ERROR_DATA_INVALID   int32 = 400002
	ERROR_BAD_WORD       int32 = 400003
	ERROR_WRONG_FILE     int32 = 400004
	ERROR_DUPLICATE_DATA int32 = 400005
	ERROR_WRONG_TIME     int32 = 400006
	ERROR_UPLOAD_S3      int32 = 400007

	ERROR_NON_CLIENT                  int32 = 403010
	ERROR_MISSING_TOKEN               int32 = 403011
	ERROR_REFRESH_TOKEN_INVALID       int32 = 403013
	ERROR_TOKEN_INVALID               int32 = 403012
	ERROR_EARLY_LOGIN                 int32 = 403014
	ERROR_UNKNOW_3RD_PROVIDER         int32 = 403015
	ERROR_OAUTH_PROVIDER_ID_INVALID   int32 = 403016
	ERROR_OAUTH_PROVIDER_ID_NOT_MATCH int32 = 403017
	ERROR_API_KEY_INVALID             int32 = 403018
	ERROR_EARLY_CODE_INVALID          int32 = 403019
	ERROR_URI_INVALID                 int32 = 403020
	ERROR_USER_WAS_BANNED             int32 = 403021
	ERROR_TOKEN_LOCATION_WAS_CHANGED  int32 = 403022
	ERROR_NEED_NEW_TOKEN              int32 = 403023
	ERROR_ACCESS_DENIED               int32 = 403024

	// OPENID
	ERROR_WRONG_CAPTCHA                  int32 = 400050
	ERROR_ACCOUNT_LOCKED                 int32 = 400051
	ERROR_USER_LOGGED_IN                 int32 = 400052
	ERROR_LOGIN_INFO_INVALID             int32 = 400053
	ERROR_ACCOUNT_ACTIVED                int32 = 400054
	ERROR_ACCOUNT_NOT_EXISTED            int32 = 400055
	ERROR_ACCOUNT_WAS_BANNED             int32 = 400056
	ERROR_REGISTER_FAIL                  int32 = 400060
	ERROR_REACH_MAX_REGISTERED           int32 = 400061
	ERROR_ACCOUNT_IN_ACTIVE              int32 = 400062
	ERROR_ROLE_TITLE_EXISTED             int32 = 400063
	ERROR_EASY_PASSWORD                  int32 = 400080
	ERROR_EASE_NEW_PASSWORD              int32 = 400081
	ERROR_WRONG_PASSWORD                 int32 = 400082
	ERROR_PASSWORD_TYPE                  int32 = 400083
	ERROR_PASSWORD_INVALID               int32 = 400084
	ERROR_PASSWORD_EXPIRED               int32 = 400085
	ERROR_PASSWORD_RESET_CODE            int32 = 400086
	ERROR_CHANGE_SAME_PASSWORD           int32 = 400087
	ERROR_PASSWORD_LENGTH                int32 = 400088
	ERROR_WRONG_USERNAME                 int32 = 400100
	ERROR_EMPTY_USERNAME_OR_PASSWORD     int32 = 400101
	ERROR_USERNAME_TAKEN                 int32 = 400102
	ERROR_USERNAME_LENGTH                int32 = 400103
	ERROR_WRONG_EMAIL_FORMAT             int32 = 400110
	ERROR_EMAIL_NOT_MATCH                int32 = 400111
	ERROR_EMAIL_TAKEN                    int32 = 400112
	ERROR_EMAIL_ACTIVED                  int32 = 400113
	ERROR_WRONG_MAIL                     int32 = 400114
	ERROR_EMAIL_EXISTED                  int32 = 400115
	ERROR_EMAIL_NOT_FOUND                int32 = 400116
	ERROR_PHONE_NUMBER_TAKEN             int32 = 400120
	ERROR_PHONE_NUMBER_INVALID           int32 = 400121
	ERROR_PHONE_NUMBER_EXISTED           int32 = 400122
	ERROR_DISPLAY_NAME_TAKEN             int32 = 400130
	ERROR_DISPLAY_NAME_FORMAT_INVALID    int32 = 400131
	ERROR_ADDRESS_INVALID                int32 = 400135
	ERROR_PERMISSION_CREATE              int32 = 400140
	ERROR_DISPLAY_NAME_CHANGE_PERMISSION int32 = 400141
	ERROR_UPLOAD_AVATAR                  int32 = 400150
	ERROR_UNKNOW_BIRD_DAY                int32 = 400155
	ERROR_WRONG_PIN_CODE                 int32 = 400156
	ERROR_EMAIL_INACTIVE                 int32 = 400157

	ERROR_FRIEND_REQUESTED         int32 = 400160
	ERROR_USER_WAS_NOT_FRIEND      int32 = 400161
	ERROR_ADD_SELF_AS_FRIEND       int32 = 400162
	ERROR_MAX_FRIENDS_REACH        int32 = 400163
	ERROR_FRIEND_MAX_FRIENDS_REACH int32 = 400164

	ERROR_AUTHORIZE                       int32 = 401000
	ERROR_UNAUTHORIZED_USER               int32 = 401001
	ERROR_USER_LOGGED_IN_SESSION_TIME_OUT int32 = 401002
	ERROR_AUTHORIZE_CODE_INVALID          int32 = 401003
	ERROR_AUTHORIZE_CODE_EXPIRED          int32 = 401004
	ERROR_SESSION_EXPIRED                 int32 = 401005
	ERROR_NOT_DEVELOPER_ACCOUNT           int32 = 401006

	ERROR_NULL_ID int32 = 400200

	// CHAT
	ERROR_DIRECT_CHAT_NOT_EXISTED      int32 = 400320
	ERROR_DIRECT_CHAT_EXISTED          int32 = 400321
	ERROR_CREATE_DIRECT_CHAT_WITH_SELF int32 = 400322
	ERROR_COULD_NOT_JOIN_DIRECT_CHAT   int32 = 400323
	ERROR_ACCEPT_DIRECT_CHAT           int32 = 400324
	ERROR_BLOCK_DIRECT_CHAT_FAIL       int32 = 400325
	ERROR_UNBLOCK_DIRECT_CHAT_FAIL     int32 = 400326
	ERROR_NOT_ABLE_TO_PUBLISH_MQTT     int32 = 400327
	ERROR_NOT_ABLE_TO_SUBSCRIBE_MQTT   int32 = 400328
	ERROR_IN_GAME_CHAT_NOT_EXISTED     int32 = 400330
	ERROR_COULD_NOT_JOIN_IN_GAME_CHAT  int32 = 400331
	ERROR_STATUS_INVALID               int32 = 400335
	ERROR_CHANNEL_NOT_EXISTED          int32 = 400340
	ERROR_COULD_NOT_JOIN_CHANNEL       int32 = 400341
	ERROR_PUBLISH_INTERVAL_NOT_ALLOWED int32 = 400342
	ERROR_SYSTEM_MAINTENANCE           int32 = 400350
	ERROR_IN_PROCESSING                int32 = 400351
	ERROR_QUERY_PARAM                  int32 = 400360
	ERROR_EXCEED_LIMIT                 int32 = 400361
	ERROR_ACTION_INVALID               int32 = 400370
	ERROR_ALREADY_CHECK_IN             int32 = 400371

	// TRANSACTION
	ERROR_ANOTHER_TRANSACTION_PROCESSING       int32 = 400380
	ERROR_TRANSACTION_INVALID                  int32 = 400381
	ERROR_TRANSACTION_ENOUGH                   int32 = 400382
	ERROR_ORDER_INVALID                        int32 = 400390
	ERROR_TRANSACTION_LOCKED                   int32 = 400391
	ERROR_TRANSACTION_SAME_CREATOR_AND_PARTNER int32 = 400392
	ERROR_TRANSACTION_TIMEOUT                  int32 = 400393

	// PAYMENT
	ERROR_EMPTY_BANK_ACOUNT_NUMBER int32 = 400400
	ERROR_EMPTY_BANK_ACOUNT_HOLDER int32 = 400401
	ERROR_WRONG_BANK_INFO          int32 = 400402
	ERROR_PAYMENT_KIND             int32 = 400403
	ERROR_AMOUNT_INVALID           int32 = 400405
	ERROR_AMOUNT_NOT_ENOUGH        int32 = 400406
	ERROR_PAYER_INVALID            int32 = 400407

	// MINI_GAME
	ERROR_INVALID_QUANTITY int32 = 400500

	// AUCTION
	ERROR_SESSION_NULL_ID        int32 = 400600
	ERROR_SESSION_LOCKED         int32 = 400601
	ERROR_PRODUCTS_LOCKED        int32 = 400602
	ERROR_BID_PRICE_EXISTS       int32 = 400603
	ERROR_PRODUCT_IS_PASSED      int32 = 400604
	ERROR_PRODUCT_IS_SOLD        int32 = 400605
	ERROR_PRODUCT_NOT_AVAILABLE  int32 = 400606
	ERROR_SESSION_OUT_OF_PRODUCT int32 = 400607
	ERROR_BILLING_EXISTS         int32 = 400608
	ERROR_PRICE_STEP_IS_NOT_SET  int32 = 400609
	ERROR_SESSION_OWNER          int32 = 400610
	ERROR_EMPTY_BUY_NOW_PRICE    int32 = 400611

	// SERVER SIDE ERROR
	ERROR_OTHER                  int32 = 500000
	ERROR_CONNECT                int32 = 500001
	ERROR_SAVE_DATA              int32 = 500002
	ERROR_RETRIEVE_DATA          int32 = 500003
	ERROR_NOT_FOUND              int32 = 500004
	ERROR_GET_USER_INFO          int32 = 500005
	ERROR_DELETE_DATA_FROM_CACHE int32 = 500006
	ERROR_NOTIFY_TOO_MUCH        int32 = 500007
	ERROR_NOT_IMPLEMENTED        int32 = 500008
	ERROR_RELOAD_POLICY          int32 = 500009
	ERROR_ADD_NEW_VERSION        int32 = 500020
	ERROR_VERSION_EXISTED        int32 = 500021
	ERROR_VERSION_NOT_EXISTED    int32 = 500022
	ERROR_VERSION_NOT_RELEASED   int32 = 500023
	ERROR_UNMARSHAL_JSON         int32 = 500030
	ERROR_ID_ALREADY_NOTIFIED    int32 = 500040
	ERROR_ID_ALREADY_SCHEDULE    int32 = 500041
	ERROR_SCHEDULE_INVALID       int32 = 500042
	ERROR_EXTRACT_FILE           int32 = 500043
	ERROR_REDIS_RUNTIME          int32 = 500044
	ERROR_CREATE_ACCOUNT         int32 = 500045
	ERROR_CREATE_TOKEN           int32 = 500046
	ERROR_RESET_PASSWORD         int32 = 500047
	ERROR_CREATE_EMAIL           int32 = 500048
	ERROR_SERVER_CONFIG          int32 = 500060
	ERROR_SERVER_LOGIC           int32 = 500070
	ERROR_CONNECT_LOCAL_SERVICE  int32 = 500080
	ERROR_INTERNAL               int32 = 500999
)

var ERROR_CODE_NAME = map[int32]string{

	202001: "ERROR_PASSWORD_LEVEL2",
	202002: "ERROR_GOOGLE_AUTH",

	0:      "ERROR_CODE_OK",
	400000: "ERROR_BAD_REQUEST",
	400001: "ERROR_BIND_DATA",
	400002: "ERROR_DATA_INVALID",
	400003: "ERROR_BAD_WORD",
	400004: "ERROR_WRONG_FILE",
	400005: "ERROR_DUPLICATE_DATA",
	400006: "ERROR_WRONG_TIME",
	400007: "ERROR_UPLOAD_S3",

	403010: "ERROR_NON_CLIENT",
	403011: "ERROR_MISSING_TOKEN",
	403012: "ERROR_TOKEN_INVALID",
	403013: "ERROR_REFRESH_TOKEN_INVALID",
	403014: "ERROR_EARLY_LOGIN",
	403015: "ERROR_UNKNOW_3RD_PROVIDER",
	403016: "ERROR_OAUTH_PROVIDER_ID_INVALID",
	403017: "ERROR_OAUTH_PROVIDER_ID_NOT_MATCH",
	403018: "ERROR_API_KEY_INVALID",
	403019: "ERROR_EARLY_CODE_INVALID",
	403020: "ERROR_URI_INVALID",
	403021: "ERROR_USER_WAS_BANNED",
	403022: "ERROR_TOKEN_LOCATION_WAS_CHANGED",
	403023: "ERROR_NEED_NEW_TOKEN",
	403024: "ERROR_ACCESS_DENIED",

	400050: "ERROR_WRONG_CAPTCHA",
	400051: "ERROR_ACCOUNT_LOCKED",
	400052: "ERROR_USER_LOGGED_IN",
	400053: "ERROR_LOGIN_INFO_INVALID",
	400054: "ERROR_ACCOUNT_ACTIVE",
	400055: "ERROR_ACCOUNT_NOT_EXISTED",
	400056: "ERROR_ACCOUNT_WAS_BANNED",
	400060: "ERROR_REGISTER_FAIL",
	400061: "ERROR_REACH_MAX_REGISTERED",
	400062: "ERROR_ACCOUNT_IN_ACTIVE",
	400063: "ERROR_ROLE_TITLE_EXISTED",
	400080: "ERROR_EASY_PASSWORD",
	400081: "ERROR_EASE_NEW_PASSWORD",
	400082: "ERROR_WRONG_PASSWORD",
	400083: "ERROR_PASSWORD_TYPE",
	400084: "ERROR_PASSWORD_INVALID",
	400085: "ERROR_PASSWORD_EXPIRED",
	400086: "ERROR_PASSWORD_RESET_CODE",
	400087: "ERROR_CHANGE_SAME_PASSWORD",
	400088: "ERROR_PASSWORD_LENGTH",
	400100: "ERROR_WRONG_USERNAME",
	400101: "ERROR_EMPTY_USERNAME_OR_PASSWORD",
	400102: "ERROR_USERNAME_TAKEN",
	400103: "ERROR_USERNAME_LENGTH",
	400110: "ERROR_WRONG_EMAIL_FORMAT",
	400111: "ERROR_EMAIL_NOT_MATCH",
	400112: "ERROR_EMAIL_TAKEN",
	400113: "ERROR_EMAIL_ACTIVE",
	400114: "ERROR_WRONG_MAIL",
	400115: "ERROR_EMAIL_EXISTED",
	400116: "ERROR_EMAIL_NOT_FOUND",
	400120: "ERROR_PHONE_NUMBER_TAKEN",
	400121: "ERROR_PHONE_NUMBER_INVALID",
	400122: "ERROR_PHONE_NUMBER_EXISTED",
	400130: "ERROR_DISPLAY_NAME_TAKEN",
	400131: "ERROR_DISPLAY_NAME_FORMAT_INVALID",
	400135: "ERROR_ADDRESS_INVALID",
	400140: "ERROR_PERMISSION_CREATE",
	400141: "ERROR_DISPLAY_NAME_CHANGE_PERMISSION",
	400150: "ERROR_UPLOAD_AVATAR",
	400155: "ERROR_UNKNOW_BIRD_DAY",
	400156: "ERROR_WRONG_PIN_CODE",
	400157: "ERROR_EMAIL_INACTIVE",

	400160: "ERROR_FRIEND_REQUESTED",
	400161: "ERROR_USER_WAS_NOT_FRIEND",
	400162: "ERROR_ADD_SELF_AS_FRIEND",
	400163: "ERROR_MAX_FRIENDS_REACH",
	400164: "ERROR_FRIEND_MAX_FRIENDS_REACH",
	401000: "ERROR_AUTHORIZE",
	401001: "ERROR_UNAUTHORIZED_USER",
	401002: "ERROR_USER_LOGGED_IN_SESSION_TIME_OUT",
	401003: "ERROR_AUTHORIZE_CODE_INVALID",
	401004: "ERROR_AUTHORIZE_CODE_EXPIRED",
	401005: "ERROR_SESSION_EXPIRED",
	401006: "ERROR_NOT_DEVELOPER_ACCOUNT",

	400200: "ERROR_NULL_ID",

	400320: "ERROR_DIRECT_CHAT_NOT_EXISTED",
	400321: "ERROR_DIRECT_CHAT_EXISTED",
	400322: "ERROR_CREATE_DIRECT_CHAT_WITH_SELF",
	400323: "ERROR_COULD_NOT_JOIN_DIRECT_CHAT",
	400324: "ERROR_ACCEPT_DIRECT_CHAT",
	400325: "ERROR_BLOCK_DIRECT_CHAT_FAIL",
	400326: "ERROR_UNBLOCK_DIRECT_CHAT_FAIL",
	400327: "ERROR_NOT_ABLE_TO_PUBLISH_MQTT",
	400328: "ERROR_NOT_ABLE_TO_SUBSCRIBE_MQTT",
	400330: "ERROR_IN_GAME_CHAT_NOT_EXISTED",
	400331: "ERROR_COULD_NOT_JOIN_IN_GAME_CHAT",
	400335: "ERROR_STATUS_INVALID",
	400340: "ERROR_CHANNEL_NOT_EXISTED",
	400341: "ERROR_COULD_NOT_JOIN_CHANNEL",
	400342: "ERROR_PUBLISH_INTERVAL_NOT_ALLOWED",
	400350: "ERROR_SYSTEM_MAINTENANCE",
	400351: "ERROR_IN_PROCESSING",
	400360: "ERROR_QUERY_PARAM",
	400361: "ERROR_EXCEED_LIMIT",
	400370: "ERROR_ACTION_INVALID",
	400371: "ERROR_ALREADY_CHECK_IN",

	400380: "ERROR_ANOTHER_TRANSACTION_PROCESSING",
	400381: "ERROR_TRANSACTION_INVALID",
	400382: "ERROR_TRANSACTION_ENOUGH",
	400390: "ERROR_ORDER_INVALID",
	400391: "ERROR_TRANSACTION_LOCKED",
	400392: "ERROR_TRANSACTION_SAME_CREATOR_AND_PARTNER",
	400393: "ERROR_TRANSACTION_TIMEOUT",

	400400: "ERROR_EMPTY_BANK_ACOUNT_NUMBER",
	400401: "ERROR_EMPTY_BANK_ACOUNT_HOLDER",
	400402: "ERROR_WRONG_BANK_INFO",
	400403: "ERROR_PAYMENT_KIND",
	400405: "ERROR_AMOUNT_INVALID",
	400406: "ERROR_AMOUNT_NOT_ENOUGH",
	400407: "ERROR_PAYER_INVALID",

	400500: "ERROR_INVALID_QUANTITY",
	400600: "ERROR_SESSION_NULL_ID",
	400601: "ERROR_SESSION_LOCKED",
	400602: "ERROR_PRODUCTS_LOCKED",
	400603: "ERROR_BID_PRICE_EXISTS",
	400604: "ERROR_PRODUCT_IS_PASSED",
	400605: "ERROR_PRODUCT_IS_SOLD",
	400606: "ERROR_PRODUCT_NOT_AVAILABLE",
	400607: "ERROR_SESSION_OUT_OF_PRODUCT",
	400608: "ERROR_BILLING_EXISTS",
	400609: "ERROR_PRICE_STEP_IS_NOT_SET",
	400610: "ERROR_SESSION_OWNER",
	400611: "ERROR_EMPTY_BUY_NOW_PRICE",

	500000: "ERROR_OTHER",
	500001: "ERROR_CONNECT",
	500002: "ERROR_SAVE_DATA",
	500003: "ERROR_RETRIEVE_DATA",
	500004: "ERROR_NOT_FOUND",
	500005: "ERROR_GET_USER_INFO",
	500006: "ERROR_DELETE_DATA_FROM_CACHE",
	500007: "ERROR_NOTIFY_TOO_MUCH",
	500008: "ERROR_NOT_IMPLEMENTED",
	500009: "ERROR_RELOAD_POLICY",
	500020: "ERROR_ADD_NEW_VERSION",
	500021: "ERROR_VERSION_EXISTED",
	500022: "ERROR_VERSION_NOT_EXISTED",
	500023: "ERROR_VERSION_NOT_RELEASED",
	500030: "ERROR_UNMARSHAL_JSON ",
	500040: "ERROR_ID_ALREADY_NOTIFIED",
	500041: "ERROR_ID_ALREADY_SCHEDULE",
	500042: "ERROR_SCHEDULE_INVALID",
	500043: "ERROR_EXTRACT_FILE",
	500044: "ERROR_REDIS_RUNTIME",
	500045: "ERROR_CREATE_ACCOUNT",
	500046: "ERROR_CREATE_TOKEN",
	500047: "ERROR_RESET_PASSWORD",
	500048: "ERROR_CREATE_EMAIL",
	500060: "ERROR_SERVER_CONFIG",
	500070: "ERROR_SERVER_LOGIC",
	500080: "ERROR_CONNECT_LOCAL_SERVICE",
	500999: "ERROR_INTERNAL",
}
