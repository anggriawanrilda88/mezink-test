package errors

// ERROR CODE TYPE: Frist number error code
// 1 -> Critical 	-> Server Issue and Blocker to Business Core flow
// 2 -> Critical 	-> Major Issue and Blocker to Business Core flow
// 3 -> Error 		-> Moderate Issue
// 4 -> Error 		-> Minor Issue
// 5 -> Warning 	-> Warning

// ==========ERROR CODE SERVICE NAME ORDER==========
// record -> 001

// ==========CTA NUMBER==========
// 00 No Action 						-> Just for information and no need to action.
// 01 Back to previous page/screen 		-> When an error occurs, it can be returned to the previous page/screen
// 02 Back to Home 						-> When an error occurs, it can be returned to the main page.
// 03 Reload/ Refresh 					-> It is recommended to reload or refresh the page.
// 04 Close 							-> It is recommended to close the page or close the application and it can be recommended to open the previous application or page again.
// 05 Relogin 							-> It is recommended to logout and login again.
// 06 Call Support 						-> To be recommended to contact our support team.

// ==========ERROR MESSAGE CODE==========
// number of error on app,  3 Digits error message code, example 4_001_00_001, 4_001_00_002

const (
	GET_LIST_RECORD_REQ_ERROR        string = "errStatus:400; errCode:4_001_00_001; message:Failed query params"
	GET_LIST_RECORD_VALIDATION_ERROR string = "errStatus:400; errCode:4_001_00_002; message:List record validation error"
	FIND_SERVICE_RECORD_ERROR        string = "errStatus:400; errCode:4_001_00_003; message:Find student error"
	CREATE_SERVICE_RECORD_ERROR      string = "errStatus:400; errCode:4_001_00_004; message:Create student error"
	CREATE_RECORD_BODY_ERROR         string = "errStatus:400; errCode:4_001_00_001; message:Failed body request"
	CREATE_RECORD_VALIDATION_ERROR   string = "errStatus:400; errCode:4_001_00_002; message:Create record validation error"
)
