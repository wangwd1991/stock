package commad

import (
	"fmt"
	"stock/db"
	"stock/server"
	"time"
)

var (
	GainRed   = "\033[1;31m+%4.2f\033[0m"
	GainGreen = "\033[1;32m%-5.2f\033[0m"
	GainWhite = "\033[1;37m%-5.2f\033[0m"
)

func Command(dbSrv db.Service) error {
	srv := server.NewService(dbSrv)
	for range time.NewTicker(5 * time.Second).C {
		stocks, err := srv.ListStocks()
		if err != nil {
			return err
		}
		if len(stocks) == 0 {
			continue
		}
		err = srv.GetStocks(stocks)
		if err != nil {
			return err
		}
		fmt.Printf("%-6s %-8s %-7s %-5s %-7s %-7s %-7s\n", "CODE", "NAME", "CUR", "GAIN", "MAX", "MIN", "YES")
		for _, st := range stocks {
			if st.Current > st.ExpectSell {
				fmt.Printf("Congratulate!!! to code %s name %s %.2f\n", st.Code, st.Name, st.Current)
			}

			formatStr := "%-6s %-4s %-7.2f " + GainWhite + " %-7.2f %-7.2f %-7.2f\n"
			if st.Gains > 0 {
				formatStr = "%-6s %-4s %-7.2f " + GainRed + " %-7.2f %-7.2f %-7.2f\n"
			} else if st.Gains < 0 {
				formatStr = "%-6s %-4s %-7.2f " + GainGreen + " %-7.2f %-7.2f %-7.2f\n"
			}
			fmt.Printf(formatStr, st.Code, st.Name, st.Current, st.Gains, st.Max, st.Min, st.Yesterday)
		}
	}
	return nil
}
