package httputils

import (
        "net/http"
        "errors"
        "regexp"
        "net/http/httputil"
        "crypto/tls"
        "bytes"
        "strings"
        "io/ioutil"
        "bufio"
        "fmt"
        "github.com/jshort/gocurl/cliutils"
)

const userAgent string = "Gocurl-client/1.0" 

var verboseOutput bool = false

var tr *http.Transport = &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
}

var client *http.Client = &http.Client{
        Transport: tr,
        CheckRedirect: redirectHandler,
}

func SubmitRequest(cliInputs *cliutils.GoCurlCli) int {
        verboseOutput = cliInputs.Verbose()
        applyColor(cliInputs.Color())
        configureTls(cliInputs)

        var retval int

        switch cliInputs.HttpVerb() {
        case "GET":
                retval = get(cliInputs)
        case "HEAD":
                retval = get(cliInputs)
        case "POST":
                retval = post(cliInputs)
        case "PUT":
                retval = put(cliInputs)
        case "DELETE":
                retval = delete(cliInputs)
        case "PATCH":
                retval = patch(cliInputs)
        }

        return retval
}

func get(cliInputs *cliutils.GoCurlCli) int {
        req, _ := http.NewRequest(cliInputs.HttpVerb(), cliInputs.Url(), nil)

        prepareRequest(req, cliInputs.HttpHeaders())

        return processRequest(req, cliInputs)
}

func post(cliInputs *cliutils.GoCurlCli) int {
        var postBody = []byte(cliInputs.PostData())
        req, _ := http.NewRequest("POST", cliInputs.Url(), bytes.NewBuffer(postBody))

        prepareRequest(req, cliInputs.HttpHeaders())

        return processRequest(req, cliInputs)
}

func put(cliInputs *cliutils.GoCurlCli) int {
        return dummyReturn()
}

func delete(cliInputs *cliutils.GoCurlCli) int {
        return dummyReturn()
}

func patch(cliInputs *cliutils.GoCurlCli) int {
        return dummyReturn()
}

func dummyReturn() int {
        fmt.Println("Not implemented, exiting...")        
        return 255
}

func prepareRequest(req *http.Request, headers []string) {
        headerMap := parseHeaderString(headers)
        req.Header.Set("User-Agent", userAgent)
        for key, value := range headerMap {
                req.Header.Set(key, value)
        }

        if verboseOutput {
                printRequest(req)
        }
}

func processRequest(req *http.Request, cliInputs *cliutils.GoCurlCli) int {
        var resp *http.Response
        var err error
        if cliInputs.Redirect() {
                resp, err = client.Do(req)
        } else {
                resp, err = tr.RoundTrip(req)
        }
        
        if err != nil {
                retval, msg := handleHttpError(err)
                fmt.Println(msg)
                return retval
        }
        defer resp.Body.Close()
        respBody, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                fmt.Printf("Error reading body: %v\n", err)
                return 1
        }

        if verboseOutput {
                printResponse(resp)
        }

        fmt.Printf("%s", respBody)
        return 0
}

func printRequest(req *http.Request) {
        reqDump, err := httputil.DumpRequestOut(req, true)
        fmt.Println(bBlue("Request:"))

        r := bufio.NewReader(bytes.NewBuffer(reqDump))
        line, isPrefix, err := r.ReadLine()
        for err == nil && !isPrefix {
                s := string(line)
                fmt.Printf("%s %s\n", bCyan(">"), bYellow(s))
                line, isPrefix, err = r.ReadLine()
        }
        fmt.Println("")
}

func printResponse(resp *http.Response) {
        fmt.Println(bBlue("Response:"))
        fmt.Printf("%s %s %s\n", bCyan("<"),
                        bRed(resp.Proto),
                        bMagenta(resp.Status))
        for header, values := range resp.Header {
                fmt.Printf("%s %s: %s\n",
                        bCyan("<"),
                        bRed(header),
                        bRed(strings.Join(values, ", ")))
        }
        fmt.Println("")
}

// Assumes the header slice was validated (1 colon per entry)
func parseHeaderString(headers []string) map[string]string {
        var headerMap = make(map[string]string)
        for _, header := range headers {
                tokens := strings.Split(header, ":")
                headerMap[tokens[0]] = tokens[1]
        }
        return headerMap
}

func configureTls(cliInputs *cliutils.GoCurlCli) {
        if ! cliInputs.SslSecure() {
                (*tr).TLSClientConfig.InsecureSkipVerify = true
        }
}

func redirectHandler(req *http.Request, via []*http.Request) error {
        if len(via) >= 10 {
                return errors.New("stopped after 10 redirects")
        }
        if verboseOutput {
                printRequest(req)
        }
        return nil
}

func handleHttpError(err error) (retval int, msg string) {
        switch {
        case strings.Contains(err.Error(), "malformed HTTP response"):
                retval = 1
                msg = "ERROR: HTTP request sent to HTTPS listener"
        case strings.Contains(err.Error(), "tls: oversized record received"):
                retval = 2
                msg = "ERROR: HTTPS request sent to HTTP listener"
        case strings.Contains(err.Error(), "getsockopt: connection refused"):
                retval = 3
                msg = "ERROR: Invalid host or port"
        case strings.Contains(err.Error(), "x509: certificate signed by unknown authority"):
                retval = 4
                msg = "ERROR: Server certificate unknown (perhaps try with -k if testing)"
        case strings.Contains(err.Error(), "no such host"):
                hostOrIpRegex := regexp.MustCompile(`lookup ([\w\.]+):`)
                retval = 5
                msg = "ERROR: No such host: " + hostOrIpRegex.FindStringSubmatch(err.Error())[1]
        default:
                retval = 255
                msg = "ERROR: " + err.Error()
        }
        return
}
