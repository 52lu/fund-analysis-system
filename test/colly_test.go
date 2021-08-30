package test

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"testing"
)

func TestXueQiu(t *testing.T) {
	collector := colly.NewCollector()
	collector.OnHTML("table[class='quote-info']", func(element *colly.HTMLElement) {
		fmt.Println(element.Text)
	})
	err := collector.Visit("https://xueqiu.com/S/SH688169")
	if err != nil {
		t.Error(err)
	}
}
