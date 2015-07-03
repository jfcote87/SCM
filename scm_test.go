package scm

import (
	"os"
	"testing"
)

func TestFileDownload(t *testing.T) {
	// Test File ID
	fid := os.Getenv("SCM_TEST_FID")

	login := os.Getenv("SCM_TEST_LOGIN")
	passwd := os.Getenv("SCM_TEST_PASSWD")
	apikey := os.Getenv("SCM_TEST_API_KEY")
	sv := NewSpringCMServiceSoap("", true)
	tk, err := sv.Authenticate(&Authenticate{UserName: login, Password: passwd, ApiKey: apikey})
	if err != nil {
		t.Fatalf("Authenticate: %v", err)
	}
	t.Logf("%s", tk.AuthenticateResult)
	fx, err := sv.DocumentGetById(&DocumentGetById{Token: tk.AuthenticateResult, DocumentId: fid, LoadExtendedMetadata: true})
	if err != nil {
		t.Fatalf("Download: %v", err)
	}
	t.Logf("Total Bytes: %d %d", fx.DocumentGetByIdResult.FileSize, fx.DocumentGetByIdResult.PDFFileSize)

	dd := DocumentDownload{}
	dd.Token = tk.AuthenticateResult
	dd.DocumentId = fid
	var z DownloadFormat = DownloadFormatPDF
	dd.Format = &z
	dd.Length = fx.DocumentGetByIdResult.PDFFileSize

	dx, err := sv.DocumentDownload(&dd)
	if err != nil {
		t.Fatalf("Download: %v", err)
	}

}
