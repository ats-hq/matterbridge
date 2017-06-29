// Code generated by protoc-gen-go.
// source: gcsystemmsgs.proto
// DO NOT EDIT!

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EGCSystemMsg int32

const (
	EGCSystemMsg_k_EGCMsgInvalid                           EGCSystemMsg = 0
	EGCSystemMsg_k_EGCMsgMulti                             EGCSystemMsg = 1
	EGCSystemMsg_k_EGCMsgGenericReply                      EGCSystemMsg = 10
	EGCSystemMsg_k_EGCMsgSystemBase                        EGCSystemMsg = 50
	EGCSystemMsg_k_EGCMsgAchievementAwarded                EGCSystemMsg = 51
	EGCSystemMsg_k_EGCMsgConCommand                        EGCSystemMsg = 52
	EGCSystemMsg_k_EGCMsgStartPlaying                      EGCSystemMsg = 53
	EGCSystemMsg_k_EGCMsgStopPlaying                       EGCSystemMsg = 54
	EGCSystemMsg_k_EGCMsgStartGameserver                   EGCSystemMsg = 55
	EGCSystemMsg_k_EGCMsgStopGameserver                    EGCSystemMsg = 56
	EGCSystemMsg_k_EGCMsgWGRequest                         EGCSystemMsg = 57
	EGCSystemMsg_k_EGCMsgWGResponse                        EGCSystemMsg = 58
	EGCSystemMsg_k_EGCMsgGetUserGameStatsSchema            EGCSystemMsg = 59
	EGCSystemMsg_k_EGCMsgGetUserGameStatsSchemaResponse    EGCSystemMsg = 60
	EGCSystemMsg_k_EGCMsgGetUserStatsDEPRECATED            EGCSystemMsg = 61
	EGCSystemMsg_k_EGCMsgGetUserStatsResponse              EGCSystemMsg = 62
	EGCSystemMsg_k_EGCMsgAppInfoUpdated                    EGCSystemMsg = 63
	EGCSystemMsg_k_EGCMsgValidateSession                   EGCSystemMsg = 64
	EGCSystemMsg_k_EGCMsgValidateSessionResponse           EGCSystemMsg = 65
	EGCSystemMsg_k_EGCMsgLookupAccountFromInput            EGCSystemMsg = 66
	EGCSystemMsg_k_EGCMsgSendHTTPRequest                   EGCSystemMsg = 67
	EGCSystemMsg_k_EGCMsgSendHTTPRequestResponse           EGCSystemMsg = 68
	EGCSystemMsg_k_EGCMsgPreTestSetup                      EGCSystemMsg = 69
	EGCSystemMsg_k_EGCMsgRecordSupportAction               EGCSystemMsg = 70
	EGCSystemMsg_k_EGCMsgGetAccountDetails_DEPRECATED      EGCSystemMsg = 71
	EGCSystemMsg_k_EGCMsgReceiveInterAppMessage            EGCSystemMsg = 73
	EGCSystemMsg_k_EGCMsgFindAccounts                      EGCSystemMsg = 74
	EGCSystemMsg_k_EGCMsgPostAlert                         EGCSystemMsg = 75
	EGCSystemMsg_k_EGCMsgGetLicenses                       EGCSystemMsg = 76
	EGCSystemMsg_k_EGCMsgGetUserStats                      EGCSystemMsg = 77
	EGCSystemMsg_k_EGCMsgGetCommands                       EGCSystemMsg = 78
	EGCSystemMsg_k_EGCMsgGetCommandsResponse               EGCSystemMsg = 79
	EGCSystemMsg_k_EGCMsgAddFreeLicense                    EGCSystemMsg = 80
	EGCSystemMsg_k_EGCMsgAddFreeLicenseResponse            EGCSystemMsg = 81
	EGCSystemMsg_k_EGCMsgGetIPLocation                     EGCSystemMsg = 82
	EGCSystemMsg_k_EGCMsgGetIPLocationResponse             EGCSystemMsg = 83
	EGCSystemMsg_k_EGCMsgSystemStatsSchema                 EGCSystemMsg = 84
	EGCSystemMsg_k_EGCMsgGetSystemStats                    EGCSystemMsg = 85
	EGCSystemMsg_k_EGCMsgGetSystemStatsResponse            EGCSystemMsg = 86
	EGCSystemMsg_k_EGCMsgSendEmail                         EGCSystemMsg = 87
	EGCSystemMsg_k_EGCMsgSendEmailResponse                 EGCSystemMsg = 88
	EGCSystemMsg_k_EGCMsgGetEmailTemplate                  EGCSystemMsg = 89
	EGCSystemMsg_k_EGCMsgGetEmailTemplateResponse          EGCSystemMsg = 90
	EGCSystemMsg_k_EGCMsgGrantGuestPass                    EGCSystemMsg = 91
	EGCSystemMsg_k_EGCMsgGrantGuestPassResponse            EGCSystemMsg = 92
	EGCSystemMsg_k_EGCMsgGetAccountDetails                 EGCSystemMsg = 93
	EGCSystemMsg_k_EGCMsgGetAccountDetailsResponse         EGCSystemMsg = 94
	EGCSystemMsg_k_EGCMsgGetPersonaNames                   EGCSystemMsg = 95
	EGCSystemMsg_k_EGCMsgGetPersonaNamesResponse           EGCSystemMsg = 96
	EGCSystemMsg_k_EGCMsgMultiplexMsg                      EGCSystemMsg = 97
	EGCSystemMsg_k_EGCMsgWebAPIRegisterInterfaces          EGCSystemMsg = 101
	EGCSystemMsg_k_EGCMsgWebAPIJobRequest                  EGCSystemMsg = 102
	EGCSystemMsg_k_EGCMsgWebAPIJobRequestHttpResponse      EGCSystemMsg = 104
	EGCSystemMsg_k_EGCMsgWebAPIJobRequestForwardResponse   EGCSystemMsg = 105
	EGCSystemMsg_k_EGCMsgMemCachedGet                      EGCSystemMsg = 200
	EGCSystemMsg_k_EGCMsgMemCachedGetResponse              EGCSystemMsg = 201
	EGCSystemMsg_k_EGCMsgMemCachedSet                      EGCSystemMsg = 202
	EGCSystemMsg_k_EGCMsgMemCachedDelete                   EGCSystemMsg = 203
	EGCSystemMsg_k_EGCMsgMemCachedStats                    EGCSystemMsg = 204
	EGCSystemMsg_k_EGCMsgMemCachedStatsResponse            EGCSystemMsg = 205
	EGCSystemMsg_k_EGCMsgSQLStats                          EGCSystemMsg = 210
	EGCSystemMsg_k_EGCMsgSQLStatsResponse                  EGCSystemMsg = 211
	EGCSystemMsg_k_EGCMsgMasterSetDirectory                EGCSystemMsg = 220
	EGCSystemMsg_k_EGCMsgMasterSetDirectoryResponse        EGCSystemMsg = 221
	EGCSystemMsg_k_EGCMsgMasterSetWebAPIRouting            EGCSystemMsg = 222
	EGCSystemMsg_k_EGCMsgMasterSetWebAPIRoutingResponse    EGCSystemMsg = 223
	EGCSystemMsg_k_EGCMsgMasterSetClientMsgRouting         EGCSystemMsg = 224
	EGCSystemMsg_k_EGCMsgMasterSetClientMsgRoutingResponse EGCSystemMsg = 225
	EGCSystemMsg_k_EGCMsgSetOptions                        EGCSystemMsg = 226
	EGCSystemMsg_k_EGCMsgSetOptionsResponse                EGCSystemMsg = 227
	EGCSystemMsg_k_EGCMsgSystemBase2                       EGCSystemMsg = 500
	EGCSystemMsg_k_EGCMsgGetPurchaseTrustStatus            EGCSystemMsg = 501
	EGCSystemMsg_k_EGCMsgGetPurchaseTrustStatusResponse    EGCSystemMsg = 502
	EGCSystemMsg_k_EGCMsgUpdateSession                     EGCSystemMsg = 503
	EGCSystemMsg_k_EGCMsgGCAccountVacStatusChange          EGCSystemMsg = 504
	EGCSystemMsg_k_EGCMsgCheckFriendship                   EGCSystemMsg = 505
	EGCSystemMsg_k_EGCMsgCheckFriendshipResponse           EGCSystemMsg = 506
	EGCSystemMsg_k_EGCMsgGetPartnerAccountLink             EGCSystemMsg = 507
	EGCSystemMsg_k_EGCMsgGetPartnerAccountLinkResponse     EGCSystemMsg = 508
	EGCSystemMsg_k_EGCMsgVSReportedSuspiciousActivity      EGCSystemMsg = 509
)

var EGCSystemMsg_name = map[int32]string{
	0:   "k_EGCMsgInvalid",
	1:   "k_EGCMsgMulti",
	10:  "k_EGCMsgGenericReply",
	50:  "k_EGCMsgSystemBase",
	51:  "k_EGCMsgAchievementAwarded",
	52:  "k_EGCMsgConCommand",
	53:  "k_EGCMsgStartPlaying",
	54:  "k_EGCMsgStopPlaying",
	55:  "k_EGCMsgStartGameserver",
	56:  "k_EGCMsgStopGameserver",
	57:  "k_EGCMsgWGRequest",
	58:  "k_EGCMsgWGResponse",
	59:  "k_EGCMsgGetUserGameStatsSchema",
	60:  "k_EGCMsgGetUserGameStatsSchemaResponse",
	61:  "k_EGCMsgGetUserStatsDEPRECATED",
	62:  "k_EGCMsgGetUserStatsResponse",
	63:  "k_EGCMsgAppInfoUpdated",
	64:  "k_EGCMsgValidateSession",
	65:  "k_EGCMsgValidateSessionResponse",
	66:  "k_EGCMsgLookupAccountFromInput",
	67:  "k_EGCMsgSendHTTPRequest",
	68:  "k_EGCMsgSendHTTPRequestResponse",
	69:  "k_EGCMsgPreTestSetup",
	70:  "k_EGCMsgRecordSupportAction",
	71:  "k_EGCMsgGetAccountDetails_DEPRECATED",
	73:  "k_EGCMsgReceiveInterAppMessage",
	74:  "k_EGCMsgFindAccounts",
	75:  "k_EGCMsgPostAlert",
	76:  "k_EGCMsgGetLicenses",
	77:  "k_EGCMsgGetUserStats",
	78:  "k_EGCMsgGetCommands",
	79:  "k_EGCMsgGetCommandsResponse",
	80:  "k_EGCMsgAddFreeLicense",
	81:  "k_EGCMsgAddFreeLicenseResponse",
	82:  "k_EGCMsgGetIPLocation",
	83:  "k_EGCMsgGetIPLocationResponse",
	84:  "k_EGCMsgSystemStatsSchema",
	85:  "k_EGCMsgGetSystemStats",
	86:  "k_EGCMsgGetSystemStatsResponse",
	87:  "k_EGCMsgSendEmail",
	88:  "k_EGCMsgSendEmailResponse",
	89:  "k_EGCMsgGetEmailTemplate",
	90:  "k_EGCMsgGetEmailTemplateResponse",
	91:  "k_EGCMsgGrantGuestPass",
	92:  "k_EGCMsgGrantGuestPassResponse",
	93:  "k_EGCMsgGetAccountDetails",
	94:  "k_EGCMsgGetAccountDetailsResponse",
	95:  "k_EGCMsgGetPersonaNames",
	96:  "k_EGCMsgGetPersonaNamesResponse",
	97:  "k_EGCMsgMultiplexMsg",
	101: "k_EGCMsgWebAPIRegisterInterfaces",
	102: "k_EGCMsgWebAPIJobRequest",
	104: "k_EGCMsgWebAPIJobRequestHttpResponse",
	105: "k_EGCMsgWebAPIJobRequestForwardResponse",
	200: "k_EGCMsgMemCachedGet",
	201: "k_EGCMsgMemCachedGetResponse",
	202: "k_EGCMsgMemCachedSet",
	203: "k_EGCMsgMemCachedDelete",
	204: "k_EGCMsgMemCachedStats",
	205: "k_EGCMsgMemCachedStatsResponse",
	210: "k_EGCMsgSQLStats",
	211: "k_EGCMsgSQLStatsResponse",
	220: "k_EGCMsgMasterSetDirectory",
	221: "k_EGCMsgMasterSetDirectoryResponse",
	222: "k_EGCMsgMasterSetWebAPIRouting",
	223: "k_EGCMsgMasterSetWebAPIRoutingResponse",
	224: "k_EGCMsgMasterSetClientMsgRouting",
	225: "k_EGCMsgMasterSetClientMsgRoutingResponse",
	226: "k_EGCMsgSetOptions",
	227: "k_EGCMsgSetOptionsResponse",
	500: "k_EGCMsgSystemBase2",
	501: "k_EGCMsgGetPurchaseTrustStatus",
	502: "k_EGCMsgGetPurchaseTrustStatusResponse",
	503: "k_EGCMsgUpdateSession",
	504: "k_EGCMsgGCAccountVacStatusChange",
	505: "k_EGCMsgCheckFriendship",
	506: "k_EGCMsgCheckFriendshipResponse",
	507: "k_EGCMsgGetPartnerAccountLink",
	508: "k_EGCMsgGetPartnerAccountLinkResponse",
	509: "k_EGCMsgVSReportedSuspiciousActivity",
}
var EGCSystemMsg_value = map[string]int32{
	"k_EGCMsgInvalid":                           0,
	"k_EGCMsgMulti":                             1,
	"k_EGCMsgGenericReply":                      10,
	"k_EGCMsgSystemBase":                        50,
	"k_EGCMsgAchievementAwarded":                51,
	"k_EGCMsgConCommand":                        52,
	"k_EGCMsgStartPlaying":                      53,
	"k_EGCMsgStopPlaying":                       54,
	"k_EGCMsgStartGameserver":                   55,
	"k_EGCMsgStopGameserver":                    56,
	"k_EGCMsgWGRequest":                         57,
	"k_EGCMsgWGResponse":                        58,
	"k_EGCMsgGetUserGameStatsSchema":            59,
	"k_EGCMsgGetUserGameStatsSchemaResponse":    60,
	"k_EGCMsgGetUserStatsDEPRECATED":            61,
	"k_EGCMsgGetUserStatsResponse":              62,
	"k_EGCMsgAppInfoUpdated":                    63,
	"k_EGCMsgValidateSession":                   64,
	"k_EGCMsgValidateSessionResponse":           65,
	"k_EGCMsgLookupAccountFromInput":            66,
	"k_EGCMsgSendHTTPRequest":                   67,
	"k_EGCMsgSendHTTPRequestResponse":           68,
	"k_EGCMsgPreTestSetup":                      69,
	"k_EGCMsgRecordSupportAction":               70,
	"k_EGCMsgGetAccountDetails_DEPRECATED":      71,
	"k_EGCMsgReceiveInterAppMessage":            73,
	"k_EGCMsgFindAccounts":                      74,
	"k_EGCMsgPostAlert":                         75,
	"k_EGCMsgGetLicenses":                       76,
	"k_EGCMsgGetUserStats":                      77,
	"k_EGCMsgGetCommands":                       78,
	"k_EGCMsgGetCommandsResponse":               79,
	"k_EGCMsgAddFreeLicense":                    80,
	"k_EGCMsgAddFreeLicenseResponse":            81,
	"k_EGCMsgGetIPLocation":                     82,
	"k_EGCMsgGetIPLocationResponse":             83,
	"k_EGCMsgSystemStatsSchema":                 84,
	"k_EGCMsgGetSystemStats":                    85,
	"k_EGCMsgGetSystemStatsResponse":            86,
	"k_EGCMsgSendEmail":                         87,
	"k_EGCMsgSendEmailResponse":                 88,
	"k_EGCMsgGetEmailTemplate":                  89,
	"k_EGCMsgGetEmailTemplateResponse":          90,
	"k_EGCMsgGrantGuestPass":                    91,
	"k_EGCMsgGrantGuestPassResponse":            92,
	"k_EGCMsgGetAccountDetails":                 93,
	"k_EGCMsgGetAccountDetailsResponse":         94,
	"k_EGCMsgGetPersonaNames":                   95,
	"k_EGCMsgGetPersonaNamesResponse":           96,
	"k_EGCMsgMultiplexMsg":                      97,
	"k_EGCMsgWebAPIRegisterInterfaces":          101,
	"k_EGCMsgWebAPIJobRequest":                  102,
	"k_EGCMsgWebAPIJobRequestHttpResponse":      104,
	"k_EGCMsgWebAPIJobRequestForwardResponse":   105,
	"k_EGCMsgMemCachedGet":                      200,
	"k_EGCMsgMemCachedGetResponse":              201,
	"k_EGCMsgMemCachedSet":                      202,
	"k_EGCMsgMemCachedDelete":                   203,
	"k_EGCMsgMemCachedStats":                    204,
	"k_EGCMsgMemCachedStatsResponse":            205,
	"k_EGCMsgSQLStats":                          210,
	"k_EGCMsgSQLStatsResponse":                  211,
	"k_EGCMsgMasterSetDirectory":                220,
	"k_EGCMsgMasterSetDirectoryResponse":        221,
	"k_EGCMsgMasterSetWebAPIRouting":            222,
	"k_EGCMsgMasterSetWebAPIRoutingResponse":    223,
	"k_EGCMsgMasterSetClientMsgRouting":         224,
	"k_EGCMsgMasterSetClientMsgRoutingResponse": 225,
	"k_EGCMsgSetOptions":                        226,
	"k_EGCMsgSetOptionsResponse":                227,
	"k_EGCMsgSystemBase2":                       500,
	"k_EGCMsgGetPurchaseTrustStatus":            501,
	"k_EGCMsgGetPurchaseTrustStatusResponse":    502,
	"k_EGCMsgUpdateSession":                     503,
	"k_EGCMsgGCAccountVacStatusChange":          504,
	"k_EGCMsgCheckFriendship":                   505,
	"k_EGCMsgCheckFriendshipResponse":           506,
	"k_EGCMsgGetPartnerAccountLink":             507,
	"k_EGCMsgGetPartnerAccountLinkResponse":     508,
	"k_EGCMsgVSReportedSuspiciousActivity":      509,
}

func (x EGCSystemMsg) Enum() *EGCSystemMsg {
	p := new(EGCSystemMsg)
	*p = x
	return p
}
func (x EGCSystemMsg) String() string {
	return proto.EnumName(EGCSystemMsg_name, int32(x))
}
func (x *EGCSystemMsg) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EGCSystemMsg_value, data, "EGCSystemMsg")
	if err != nil {
		return err
	}
	*x = EGCSystemMsg(value)
	return nil
}
func (EGCSystemMsg) EnumDescriptor() ([]byte, []int) { return system_fileDescriptor0, []int{0} }

type ESOMsg int32

const (
	ESOMsg_k_ESOMsg_Create                   ESOMsg = 21
	ESOMsg_k_ESOMsg_Update                   ESOMsg = 22
	ESOMsg_k_ESOMsg_Destroy                  ESOMsg = 23
	ESOMsg_k_ESOMsg_CacheSubscribed          ESOMsg = 24
	ESOMsg_k_ESOMsg_CacheUnsubscribed        ESOMsg = 25
	ESOMsg_k_ESOMsg_UpdateMultiple           ESOMsg = 26
	ESOMsg_k_ESOMsg_CacheSubscriptionCheck   ESOMsg = 27
	ESOMsg_k_ESOMsg_CacheSubscriptionRefresh ESOMsg = 28
	ESOMsg_k_ESOMsg_CacheSubscribedUpToDate  ESOMsg = 29
)

var ESOMsg_name = map[int32]string{
	21: "k_ESOMsg_Create",
	22: "k_ESOMsg_Update",
	23: "k_ESOMsg_Destroy",
	24: "k_ESOMsg_CacheSubscribed",
	25: "k_ESOMsg_CacheUnsubscribed",
	26: "k_ESOMsg_UpdateMultiple",
	27: "k_ESOMsg_CacheSubscriptionCheck",
	28: "k_ESOMsg_CacheSubscriptionRefresh",
	29: "k_ESOMsg_CacheSubscribedUpToDate",
}
var ESOMsg_value = map[string]int32{
	"k_ESOMsg_Create":                   21,
	"k_ESOMsg_Update":                   22,
	"k_ESOMsg_Destroy":                  23,
	"k_ESOMsg_CacheSubscribed":          24,
	"k_ESOMsg_CacheUnsubscribed":        25,
	"k_ESOMsg_UpdateMultiple":           26,
	"k_ESOMsg_CacheSubscriptionCheck":   27,
	"k_ESOMsg_CacheSubscriptionRefresh": 28,
	"k_ESOMsg_CacheSubscribedUpToDate":  29,
}

func (x ESOMsg) Enum() *ESOMsg {
	p := new(ESOMsg)
	*p = x
	return p
}
func (x ESOMsg) String() string {
	return proto.EnumName(ESOMsg_name, int32(x))
}
func (x *ESOMsg) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ESOMsg_value, data, "ESOMsg")
	if err != nil {
		return err
	}
	*x = ESOMsg(value)
	return nil
}
func (ESOMsg) EnumDescriptor() ([]byte, []int) { return system_fileDescriptor0, []int{1} }

type EGCBaseClientMsg int32

const (
	EGCBaseClientMsg_k_EMsgGCPingRequest   EGCBaseClientMsg = 3001
	EGCBaseClientMsg_k_EMsgGCPingResponse  EGCBaseClientMsg = 3002
	EGCBaseClientMsg_k_EMsgGCClientWelcome EGCBaseClientMsg = 4004
	EGCBaseClientMsg_k_EMsgGCServerWelcome EGCBaseClientMsg = 4005
	EGCBaseClientMsg_k_EMsgGCClientHello   EGCBaseClientMsg = 4006
	EGCBaseClientMsg_k_EMsgGCServerHello   EGCBaseClientMsg = 4007
	EGCBaseClientMsg_k_EMsgGCClientGoodbye EGCBaseClientMsg = 4008
	EGCBaseClientMsg_k_EMsgGCServerGoodbye EGCBaseClientMsg = 4009
)

var EGCBaseClientMsg_name = map[int32]string{
	3001: "k_EMsgGCPingRequest",
	3002: "k_EMsgGCPingResponse",
	4004: "k_EMsgGCClientWelcome",
	4005: "k_EMsgGCServerWelcome",
	4006: "k_EMsgGCClientHello",
	4007: "k_EMsgGCServerHello",
	4008: "k_EMsgGCClientGoodbye",
	4009: "k_EMsgGCServerGoodbye",
}
var EGCBaseClientMsg_value = map[string]int32{
	"k_EMsgGCPingRequest":   3001,
	"k_EMsgGCPingResponse":  3002,
	"k_EMsgGCClientWelcome": 4004,
	"k_EMsgGCServerWelcome": 4005,
	"k_EMsgGCClientHello":   4006,
	"k_EMsgGCServerHello":   4007,
	"k_EMsgGCClientGoodbye": 4008,
	"k_EMsgGCServerGoodbye": 4009,
}

func (x EGCBaseClientMsg) Enum() *EGCBaseClientMsg {
	p := new(EGCBaseClientMsg)
	*p = x
	return p
}
func (x EGCBaseClientMsg) String() string {
	return proto.EnumName(EGCBaseClientMsg_name, int32(x))
}
func (x *EGCBaseClientMsg) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EGCBaseClientMsg_value, data, "EGCBaseClientMsg")
	if err != nil {
		return err
	}
	*x = EGCBaseClientMsg(value)
	return nil
}
func (EGCBaseClientMsg) EnumDescriptor() ([]byte, []int) { return system_fileDescriptor0, []int{2} }

type EGCToGCMsg int32

const (
	EGCToGCMsg_k_EGCToGCMsgMasterAck                   EGCToGCMsg = 150
	EGCToGCMsg_k_EGCToGCMsgMasterAckResponse           EGCToGCMsg = 151
	EGCToGCMsg_k_EGCToGCMsgRouted                      EGCToGCMsg = 152
	EGCToGCMsg_k_EGCToGCMsgRoutedReply                 EGCToGCMsg = 153
	EGCToGCMsg_k_EMsgGCUpdateSubGCSessionInfo          EGCToGCMsg = 154
	EGCToGCMsg_k_EMsgGCRequestSubGCSessionInfo         EGCToGCMsg = 155
	EGCToGCMsg_k_EMsgGCRequestSubGCSessionInfoResponse EGCToGCMsg = 156
	EGCToGCMsg_k_EGCToGCMsgMasterStartupComplete       EGCToGCMsg = 157
	EGCToGCMsg_k_EMsgGCToGCSOCacheSubscribe            EGCToGCMsg = 158
	EGCToGCMsg_k_EMsgGCToGCSOCacheUnsubscribe          EGCToGCMsg = 159
)

var EGCToGCMsg_name = map[int32]string{
	150: "k_EGCToGCMsgMasterAck",
	151: "k_EGCToGCMsgMasterAckResponse",
	152: "k_EGCToGCMsgRouted",
	153: "k_EGCToGCMsgRoutedReply",
	154: "k_EMsgGCUpdateSubGCSessionInfo",
	155: "k_EMsgGCRequestSubGCSessionInfo",
	156: "k_EMsgGCRequestSubGCSessionInfoResponse",
	157: "k_EGCToGCMsgMasterStartupComplete",
	158: "k_EMsgGCToGCSOCacheSubscribe",
	159: "k_EMsgGCToGCSOCacheUnsubscribe",
}
var EGCToGCMsg_value = map[string]int32{
	"k_EGCToGCMsgMasterAck":                   150,
	"k_EGCToGCMsgMasterAckResponse":           151,
	"k_EGCToGCMsgRouted":                      152,
	"k_EGCToGCMsgRoutedReply":                 153,
	"k_EMsgGCUpdateSubGCSessionInfo":          154,
	"k_EMsgGCRequestSubGCSessionInfo":         155,
	"k_EMsgGCRequestSubGCSessionInfoResponse": 156,
	"k_EGCToGCMsgMasterStartupComplete":       157,
	"k_EMsgGCToGCSOCacheSubscribe":            158,
	"k_EMsgGCToGCSOCacheUnsubscribe":          159,
}

func (x EGCToGCMsg) Enum() *EGCToGCMsg {
	p := new(EGCToGCMsg)
	*p = x
	return p
}
func (x EGCToGCMsg) String() string {
	return proto.EnumName(EGCToGCMsg_name, int32(x))
}
func (x *EGCToGCMsg) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EGCToGCMsg_value, data, "EGCToGCMsg")
	if err != nil {
		return err
	}
	*x = EGCToGCMsg(value)
	return nil
}
func (EGCToGCMsg) EnumDescriptor() ([]byte, []int) { return system_fileDescriptor0, []int{3} }

func init() {
	proto.RegisterEnum("EGCSystemMsg", EGCSystemMsg_name, EGCSystemMsg_value)
	proto.RegisterEnum("ESOMsg", ESOMsg_name, ESOMsg_value)
	proto.RegisterEnum("EGCBaseClientMsg", EGCBaseClientMsg_name, EGCBaseClientMsg_value)
	proto.RegisterEnum("EGCToGCMsg", EGCToGCMsg_name, EGCToGCMsg_value)
}

var system_fileDescriptor0 = []byte{
	// 1379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x57, 0x49, 0x73, 0x1b, 0x45,
	0x14, 0xce, 0x44, 0x05, 0x87, 0x2e, 0x28, 0x5e, 0x3a, 0x89, 0xed, 0x24, 0x4e, 0x94, 0x84, 0x2c,
	0xc4, 0x50, 0x39, 0x84, 0x7d, 0x47, 0x91, 0x64, 0x59, 0xc1, 0x8e, 0x15, 0x49, 0xb6, 0xd9, 0xcd,
	0x78, 0xd4, 0xb6, 0xa6, 0x2c, 0x4d, 0x0f, 0xdd, 0x3d, 0x06, 0xdd, 0xf8, 0x13, 0xac, 0x61, 0xb9,
	0xb0, 0xfe, 0x04, 0xf8, 0x05, 0xac, 0x17, 0xb8, 0xb2, 0x73, 0x84, 0x23, 0xfb, 0x52, 0xc5, 0x9b,
	0xad, 0xa7, 0x47, 0x8b, 0xb9, 0x8d, 0xfa, 0x7b, 0x7b, 0x7f, 0xef, 0xf5, 0x13, 0xa1, 0x5b, 0x8e,
	0x1c, 0x48, 0xc5, 0xfa, 0x7d, 0xb9, 0x25, 0xcf, 0xfb, 0x82, 0x2b, 0x3e, 0x77, 0xf5, 0x00, 0xb9,
	0xae, 0x5a, 0x2b, 0xb7, 0xa2, 0xf3, 0x25, 0xb9, 0x45, 0xf7, 0x93, 0x1b, 0xb6, 0xd7, 0xf1, 0x04,
	0xbf, 0xeb, 0xde, 0x8e, 0xdd, 0x73, 0x3b, 0xb0, 0x87, 0xee, 0x23, 0xd7, 0xa7, 0x87, 0x4b, 0x41,
	0x4f, 0xb9, 0x60, 0xd1, 0x19, 0x72, 0x20, 0x3d, 0xaa, 0x31, 0x8f, 0x09, 0xd7, 0x69, 0x32, 0xbf,
	0x37, 0x00, 0x42, 0xa7, 0x08, 0x4d, 0x91, 0xd8, 0xec, 0x45, 0x5b, 0x32, 0xb8, 0x40, 0x8f, 0x91,
	0xc3, 0xe9, 0x79, 0xc9, 0xe9, 0xba, 0x6c, 0x87, 0xf5, 0x99, 0xa7, 0x4a, 0xcf, 0xda, 0xa2, 0xc3,
	0x3a, 0x70, 0xab, 0xa9, 0x57, 0xe6, 0x5e, 0x99, 0xf7, 0xfb, 0xb6, 0xd7, 0x81, 0xdb, 0x4c, 0x4f,
	0x2d, 0x65, 0x0b, 0xd5, 0xe8, 0xd9, 0x03, 0xd7, 0xdb, 0x82, 0xdb, 0xe9, 0x34, 0xd9, 0x9f, 0x21,
	0xdc, 0x4f, 0x81, 0x3b, 0xe8, 0x11, 0x32, 0x9d, 0x53, 0xa9, 0xd9, 0x7d, 0x26, 0x99, 0xd8, 0x61,
	0x02, 0xee, 0xa4, 0x87, 0xc9, 0x94, 0xa9, 0x65, 0x60, 0x77, 0xd1, 0x83, 0x64, 0x5f, 0x8a, 0xad,
	0xd5, 0x9a, 0xec, 0x99, 0x80, 0x49, 0x05, 0x77, 0x9b, 0xa1, 0x85, 0xc7, 0xd2, 0xe7, 0x1e, 0xa6,
	0x74, 0x0f, 0x3d, 0x49, 0x8e, 0x65, 0x45, 0x50, 0x2b, 0x68, 0x26, 0xb4, 0x86, 0x2e, 0x95, 0x6c,
	0x39, 0x5d, 0xd6, 0xb7, 0xe1, 0x5e, 0x3a, 0x47, 0xce, 0xec, 0x2e, 0xa3, 0xed, 0xdd, 0x37, 0xc6,
	0x5e, 0x24, 0x57, 0xa9, 0x36, 0x9a, 0xd5, 0x72, 0xa9, 0x5d, 0xad, 0xc0, 0xfd, 0xf4, 0x38, 0x99,
	0x1d, 0x27, 0xa3, 0xad, 0x3c, 0x60, 0x26, 0x58, 0xf2, 0xfd, 0xba, 0xb7, 0xc9, 0x57, 0xfc, 0x8e,
	0xad, 0xb0, 0xc8, 0x0f, 0x9a, 0x95, 0x59, 0x0d, 0x2f, 0x17, 0x8f, 0x5b, 0x4c, 0x4a, 0x97, 0x7b,
	0xf0, 0x10, 0xbd, 0x91, 0x14, 0x27, 0x80, 0xda, 0x7a, 0xc9, 0x8c, 0x71, 0x91, 0xf3, 0xed, 0xc0,
	0x2f, 0x39, 0x0e, 0x0f, 0x3c, 0x35, 0x2f, 0x78, 0xbf, 0xee, 0xf9, 0x81, 0x82, 0x8b, 0xb9, 0xfa,
	0x33, 0xaf, 0xb3, 0xd0, 0x6e, 0x37, 0xd2, 0x62, 0x96, 0x4d, 0x2f, 0x43, 0xa0, 0xf6, 0x52, 0x31,
	0x2f, 0xbd, 0x21, 0x58, 0x1b, 0xc1, 0x16, 0x53, 0x81, 0x0f, 0x55, 0x5a, 0x24, 0x47, 0x52, 0xa4,
	0xc9, 0x1c, 0x2e, 0x3a, 0xad, 0xc0, 0xf7, 0xb9, 0x50, 0x25, 0x47, 0x85, 0x59, 0xcc, 0xd3, 0x9b,
	0xc8, 0x29, 0xa3, 0x40, 0x49, 0x74, 0x15, 0xa6, 0x6c, 0xb7, 0x27, 0xd7, 0x8d, 0x52, 0xd6, 0xcc,
	0x54, 0xd0, 0x14, 0x73, 0x77, 0x58, 0xdd, 0x53, 0x4c, 0x60, 0xd1, 0x96, 0x30, 0x6d, 0x7b, 0x8b,
	0x41, 0xdd, 0x0c, 0x64, 0xde, 0xf5, 0x3a, 0x89, 0x39, 0x09, 0x97, 0x4c, 0xae, 0x34, 0xb8, 0x54,
	0xa5, 0x1e, 0x13, 0x0a, 0x1e, 0x36, 0x49, 0x89, 0xee, 0x17, 0x5d, 0x87, 0x61, 0x46, 0x12, 0x16,
	0xf3, 0x1d, 0x93, 0x5d, 0x1c, 0x2c, 0x0d, 0xa9, 0x24, 0xcc, 0x97, 0x70, 0xd9, 0xcc, 0xd5, 0x00,
	0x74, 0x99, 0x96, 0x73, 0x57, 0xdd, 0xe9, 0xcc, 0x0b, 0xc6, 0x12, 0x87, 0xd0, 0x30, 0xb3, 0xcb,
	0x63, 0x5a, 0xff, 0x0a, 0x3d, 0x44, 0x0e, 0x1a, 0x0e, 0xea, 0x8d, 0x45, 0xee, 0xd8, 0x51, 0x19,
	0x9b, 0xf4, 0x04, 0x39, 0x3a, 0x16, 0xd2, 0xda, 0x2d, 0x7a, 0x94, 0x1c, 0xca, 0x77, 0xba, 0xc9,
	0xfc, 0xb6, 0x19, 0x1c, 0x5a, 0x30, 0x24, 0x60, 0x65, 0x88, 0xe9, 0x06, 0xa6, 0xcd, 0xaf, 0x9a,
	0x05, 0x0e, 0x89, 0x52, 0xed, 0xe3, 0x0d, 0xc2, 0x5a, 0xce, 0x6b, 0x7a, 0xac, 0xb5, 0x1e, 0xa1,
	0xb3, 0x64, 0xc6, 0xb0, 0x1c, 0xa1, 0x6d, 0xd6, 0xf7, 0x7b, 0x48, 0x66, 0x78, 0x94, 0x9e, 0x22,
	0xc7, 0x27, 0xa1, 0xda, 0xc6, 0x63, 0xb9, 0xc8, 0x85, 0xed, 0xa9, 0x5a, 0xc8, 0xce, 0x86, 0x2d,
	0x25, 0x3c, 0x9e, 0x8b, 0x3c, 0x87, 0x69, 0xfd, 0x27, 0xcc, 0x10, 0x47, 0x28, 0x08, 0x4f, 0xd2,
	0xd3, 0xe4, 0xc4, 0x44, 0x58, 0x5b, 0x79, 0xca, 0xec, 0x22, 0x14, 0x6b, 0x30, 0x21, 0xb9, 0x67,
	0x5f, 0x0e, 0xc7, 0x15, 0xac, 0x9b, 0x5d, 0x34, 0x04, 0x6a, 0x0b, 0x4f, 0x9b, 0x94, 0x8b, 0xe6,
	0xb6, 0xdf, 0x63, 0xcf, 0xe1, 0x37, 0xd8, 0x66, 0x1d, 0xd6, 0xd8, 0x46, 0xa9, 0x51, 0x6f, 0xb2,
	0x2d, 0x17, 0x2f, 0x41, 0x44, 0x1d, 0xb0, 0x69, 0x3b, 0xe8, 0x84, 0x99, 0xb5, 0x8c, 0xa5, 0x2e,
	0xf1, 0x8d, 0xb4, 0x91, 0x37, 0xcd, 0x46, 0x1b, 0x46, 0x17, 0x94, 0xf2, 0x75, 0x1c, 0x5d, 0x7a,
	0x33, 0x39, 0x3b, 0x49, 0x72, 0x9e, 0x8b, 0xf0, 0x05, 0xd0, 0xc2, 0x2e, 0x72, 0x32, 0x0b, 0x9a,
	0xf5, 0xcb, 0x36, 0xd2, 0xa9, 0x83, 0x29, 0xc2, 0x47, 0x16, 0x72, 0x72, 0x76, 0x1c, 0xa4, 0x95,
	0x3f, 0xb6, 0xc6, 0x6a, 0xe3, 0xe8, 0x80, 0x4f, 0x2c, 0xcc, 0x66, 0x7a, 0x04, 0xaa, 0xb0, 0x1e,
	0x43, 0x62, 0x7c, 0x6a, 0x61, 0xb5, 0xa7, 0x46, 0x15, 0x23, 0xb6, 0x7e, 0x66, 0x61, 0xb5, 0x8f,
	0x8d, 0x07, 0xb5, 0xeb, 0xcf, 0x2d, 0xe4, 0x2b, 0x68, 0x62, 0x5e, 0x59, 0x8c, 0x75, 0xbf, 0xb0,
	0x90, 0x0c, 0x33, 0xc3, 0xc7, 0x5a, 0xeb, 0x4b, 0x0b, 0x7b, 0x5c, 0x3f, 0x8b, 0x4b, 0x76, 0x78,
	0x03, 0x18, 0x6d, 0xc5, 0x15, 0xcc, 0x51, 0x5c, 0x0c, 0xe0, 0x2b, 0x8b, 0x9e, 0x25, 0x27, 0x27,
	0x0b, 0x68, 0x4b, 0x5f, 0xe7, 0x83, 0x4c, 0x05, 0x93, 0xcb, 0xe5, 0x81, 0x0a, 0x5f, 0xc6, 0x6f,
	0x2c, 0xbc, 0x8a, 0x33, 0xbb, 0x0b, 0x69, 0x8b, 0xdf, 0x5a, 0xf4, 0x4c, 0x46, 0x54, 0x2d, 0x5c,
	0xee, 0xb9, 0xf8, 0x6c, 0x87, 0x23, 0x33, 0x31, 0xfa, 0x9d, 0x45, 0xcf, 0x93, 0x73, 0xff, 0x2b,
	0xa7, 0xed, 0x7e, 0x6f, 0xe1, 0xc0, 0xcb, 0x56, 0x04, 0xa6, 0x96, 0xfd, 0x70, 0xae, 0x48, 0xf8,
	0x21, 0x57, 0x8c, 0x0c, 0xd0, 0x9a, 0x3f, 0x86, 0x6b, 0xc7, 0xfe, 0xd1, 0xe5, 0xe2, 0x02, 0xfc,
	0x52, 0x30, 0xb3, 0x0f, 0x1b, 0x22, 0x10, 0x4e, 0x17, 0xa1, 0xb6, 0x08, 0xf0, 0xe9, 0xc0, 0x9a,
	0x07, 0x12, 0x7e, 0x2d, 0x98, 0xd9, 0x8f, 0x17, 0xd2, 0xbe, 0x7e, 0x2b, 0xe0, 0x14, 0xd0, 0xc3,
	0x31, 0x7e, 0x40, 0xd3, 0x97, 0xf2, 0xf7, 0x02, 0xb6, 0x70, 0x36, 0x47, 0xca, 0x49, 0x07, 0xaf,
	0xda, 0x4e, 0x6c, 0xa4, 0xdc, 0xb5, 0x3d, 0x7c, 0x3c, 0xfe, 0x28, 0x98, 0x94, 0x2b, 0x77, 0x99,
	0xb3, 0x3d, 0x2f, 0xb0, 0x28, 0x1d, 0xd9, 0x75, 0x7d, 0xf8, 0xb3, 0x80, 0x4d, 0x58, 0x9c, 0x80,
	0xea, 0x30, 0xfe, 0x2a, 0xe0, 0xc0, 0x31, 0x07, 0x71, 0x03, 0xd7, 0x19, 0x5c, 0xb7, 0x12, 0x97,
	0x8b, 0xae, 0xb7, 0x0d, 0x7f, 0x17, 0x70, 0xc9, 0x38, 0xbd, 0xab, 0x8c, 0xb6, 0xf7, 0x4f, 0x81,
	0x9e, 0xcb, 0xda, 0x76, 0xb5, 0x85, 0x4b, 0x1b, 0xbe, 0x9d, 0x48, 0xe6, 0x40, 0xfa, 0xae, 0xe3,
	0xf2, 0x40, 0x86, 0xef, 0xe8, 0x8e, 0xab, 0x06, 0xf0, 0x6f, 0x61, 0xee, 0x85, 0xbd, 0xe4, 0xda,
	0x6a, 0x6b, 0x39, 0xdb, 0x0b, 0xa3, 0xef, 0xf5, 0xb2, 0x60, 0xe1, 0x34, 0x3d, 0x98, 0x3b, 0x8c,
	0x4b, 0x04, 0x53, 0xf4, 0x40, 0xd4, 0x06, 0xf1, 0x61, 0x05, 0x3b, 0x5c, 0xf0, 0x01, 0x4c, 0x27,
	0xa3, 0x24, 0xd1, 0x0f, 0xfb, 0xa7, 0x15, 0x6c, 0x48, 0x47, 0xb8, 0x1b, 0xb8, 0x96, 0xcc, 0x24,
	0xbb, 0xa1, 0x81, 0xae, 0x78, 0x32, 0xc3, 0x0f, 0x25, 0xa3, 0xd0, 0x74, 0x94, 0xce, 0x33, 0x38,
	0x9c, 0x8c, 0xc2, 0x51, 0xd3, 0x11, 0x7b, 0xa2, 0xc2, 0xc2, 0x91, 0x64, 0xe6, 0x4e, 0x10, 0x6a,
	0xb2, 0x4d, 0xc1, 0x64, 0x17, 0x66, 0x93, 0xb9, 0x38, 0x36, 0xcc, 0x15, 0xbf, 0xcd, 0x2b, 0x61,
	0x8a, 0x47, 0xe7, 0x7e, 0xb2, 0x08, 0x60, 0x05, 0x43, 0xee, 0x69, 0x9a, 0x27, 0xd4, 0x8c, 0x08,
	0xd1, 0x88, 0xf8, 0x1e, 0xcf, 0xc9, 0x0f, 0xa6, 0x93, 0x99, 0x64, 0x20, 0xc9, 0x65, 0x7c, 0x38,
	0x9d, 0x70, 0x2c, 0x82, 0x62, 0x4b, 0x6b, 0xac, 0xe7, 0xf0, 0x3e, 0x83, 0xb7, 0x8a, 0x26, 0xd6,
	0x8a, 0x16, 0xd4, 0x14, 0x7b, 0xbb, 0x68, 0x3a, 0x8b, 0xf5, 0x16, 0x58, 0xaf, 0xc7, 0xe1, 0x9d,
	0x1c, 0x12, 0x6b, 0xc5, 0xc8, 0xbb, 0xc5, 0x51, 0x5f, 0x35, 0xce, 0x3b, 0x1b, 0x03, 0x06, 0xef,
	0x8d, 0xf1, 0x95, 0x62, 0xef, 0x17, 0xe7, 0x7e, 0xde, 0x4b, 0x08, 0x66, 0xdb, 0xe6, 0x11, 0x67,
	0x74, 0x5b, 0x24, 0xbf, 0xe3, 0x86, 0x2f, 0x61, 0x91, 0x5f, 0xb4, 0x34, 0x57, 0x87, 0x31, 0x9d,
	0xf2, 0x4b, 0x59, 0xf3, 0x27, 0x32, 0xe1, 0x78, 0xc0, 0x3b, 0x7e, 0x39, 0x9b, 0xcf, 0x39, 0x20,
	0xfe, 0x57, 0xf1, 0x4a, 0x3a, 0xdd, 0xa2, 0x08, 0x93, 0x6e, 0x0c, 0x36, 0xc2, 0x60, 0xa3, 0x96,
	0x0c, 0x97, 0x5c, 0x78, 0xd5, 0x4a, 0x3a, 0x2a, 0x12, 0x4a, 0xea, 0x3f, 0x22, 0x75, 0xd5, 0xa2,
	0xb7, 0x44, 0xcf, 0xd1, 0x6e, 0x52, 0x3a, 0xde, 0xd7, 0xb2, 0x21, 0x98, 0xcb, 0x29, 0xfa, 0x5b,
	0x11, 0xf8, 0xb8, 0x92, 0xf9, 0xd1, 0x03, 0xf2, 0x7a, 0xfa, 0x38, 0x45, 0x56, 0x43, 0xd1, 0xd6,
	0x72, 0x9e, 0x3f, 0xf0, 0x46, 0x2e, 0x07, 0x43, 0xc4, 0xe0, 0x3a, 0xbc, 0x69, 0x5d, 0xbc, 0x66,
	0xc1, 0x7a, 0xde, 0xda, 0xf3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xd8, 0x05, 0xd1, 0xae,
	0x0d, 0x00, 0x00,
}