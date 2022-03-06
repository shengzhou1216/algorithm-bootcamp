package week2
import (
	"strings"
	"fmt"
	"strconv"
)

func subdomainVisits(cpdomains []string) []string {
    // 统计域名被访问的次数
    h := map[string]int{}
    for _,cpdomain := range cpdomains {
        countstr := strings.Split(cpdomain," ")[0]
        domain := strings.Split(cpdomain," ")[1]
        count,_ := strconv.Atoi(countstr)
        domains := strings.Split(domain,".")
        t := domains[len(domains)-1]
        h[t] += count
        for i := len(domains) - 2; i >= 0; i-- {
            t = domains[i] + "." +  t
            h[t] += count
        }
    }
    res := []string{}
    for k,v := range h {
        res = append(res,fmt.Sprintf("%d %s",v,k))
    }
    return res
}