package winapi

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

type IID GUID
type CLSID GUID
type REFIID *IID
type REFCLSID *CLSID
