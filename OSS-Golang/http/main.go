package main

import (
	"cmp"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"slices"
	"sort"
	"time"
)

// type Item struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// }

// func Sample_API1() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	//リクエストするものを作成
// 	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	//リクエスト送信
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		panic(fmt.Errorf("bad status: %s", resp.Status))
// 	}

// 	var out1 map[string]any
// 	if err := json.NewDecoder(resp.Body).Decode(&out1); err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("OK: %+v\n", out1)
// }

// // サンプルAPIのfetch関数
// func Sample_API2() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	//メソッド作成
// 	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://jsonplaceholder.typicode.com/todos/2", nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	//リクエスト実行
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()
// 	//fetchしたものをobjectに格納
// 	var out2 map[string]any
// 	if err := json.NewDecoder(resp.Body).Decode(&out2); err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("OK: %+v\n", out2)
// }

type GoID int64
type GoroutineSummary struct {
	ID         GoID
	Name       string
	TotalTime  time.Duration
	ExecTime   time.Duration
	RangeTime  map[string]time.Duration // 特殊レンジ（GC等）
	statsCache map[string]time.Duration // NonOverlappingStats のダミー
}

func (g *GoroutineSummary) NonOverlappingStats() map[string]time.Duration {
	// 実務では重複しない統計を返す。ここではダミー。
	if g.statsCache != nil {
		return g.statsCache
	}
	g.statsCache = map[string]time.Duration{
		"Execution time":  g.ExecTime,
		"Sched wait time": g.TotalTime - g.ExecTime,
	}
	return g.statsCache
}

// ダミーデータ、add(~~~)をするとデータ作成できる関数作成
func fakeSummaries() map[GoID]*GoroutineSummary {
	m := map[GoID]*GoroutineSummary{}
	now := time.Second //単位設定
	add := func(id int, name string, total, exec time.Duration, gc time.Duration) {
		m[GoID(id)] = &GoroutineSummary{
			ID:        GoID(id),
			Name:      name,
			TotalTime: total,
			ExecTime:  exec,
			RangeTime: map[string]time.Duration{
				"GC assist time": gc,
			},
		}
	}
	add(1, "main.func1", 7*now, 5*now, 1*now)
	add(2, "main.func1", 4*now, 2*now, 0)
	add(3, "worker", 9*now, 6*now, 2*now)
	add(4, "worker", 3*now, 1*now, 0)
	add(5, "listener", 5*now, 4*now, 0)
	return m
}

// サマリーから集計する
func GoroutinesHandlerFunc(summaries map[GoID]*GoroutineSummary) http.HandlerFunc {
	type group struct {
		Name     string
		N        int
		ExecTime time.Duration
	}
	tpl := template.Must(template.New("list").Parse(`
	<!doctype html><meta charset="utf-8">
	<h1>Goroutines</h1>
	<table border=1 cellpadding=6>
	<tr><th>Start location</th><th>Count</th><th>Total execution time</th></tr>
	{{range .}}
	  <tr>
	    <td><a href="/goroutine?name={{.Name}}"><code>{{.Name}}</code></a></td>
	    <td>{{.N}}</td>
	    <td>{{.ExecTime}}</td>
	  </tr>
	{{end}}
	</table>
	`))
	return func(w http.ResponseWriter, r *http.Request) {
		byName := map[string]group{}
		for _, s := range summaries {
			g := byName[s.Name] //なかったら0値、あればそのデータを返す
			g.Name = s.Name
			g.N++
			g.ExecTime += g.ExecTime
			byName[s.Name] = g //すでにbynameのkeyには値が入ってるからあとはvalueを入れてあげればいい話
		}
		//sortする
		var gs []group
		for _, g := range byName {
			gs = append(gs, g)
		}
		slices.SortFunc(gs, func(a, b group) int { return cmp.Compare(b.ExecTime, a.ExecTime) })
		if err := tpl.Execute(w, gs); err != nil {
			log.Println(err)
		}
	}
}

// 詳細
func GoroutinesHandler(summaries map[GoID]*GoroutineSummary) http.HandlerFunc {
	funcs := template.FuncMap{
		"percent": func(div, all time.Duration) string {
			if all == 0 {
				return ""
			}
			return fmt.Sprintf("(%.1f%%)", float64(div)/float64(all)*100)
		},
		"bar": func(div, max time.Duration) template.HTMLAttr {
			w := "0%"
			if max > 0 {
				w = fmt.Sprintf("%.2f%%", float64(div)/float64(max)*100)
			}
			return template.HTMLAttr(`style="height:10px;display:inline-block;background:#d7191c;width:` + w + `;"`)
		},
	}
	tpl := template.Must(template.New("detail").Funcs(funcs).Parse(`
	<!doctype html><meta charset="utf-8">
	<h1>Goroutines: {{.Name}}</h1>
	<p>Count: {{.N}} / ExecTime: {{.ExecTimePercent}}</p>
	<table border=1 cellpadding=6>
	<tr><th>ID</th><th>Total</th><th>Exec</th><th>Sched wait</th></tr>
	{{range .Goroutines}}
	  <tr>
	    <td><a href="/trace?goid={{.ID}}">{{.ID}}</a></td>
	    <td>{{.TotalTime}}</td>
	    <td>{{index .NonOverlappingStats "Execution time"}} <span {{bar (index .NonOverlappingStats "Execution time") $.MaxTotal}}></span></td>
	    <td>{{index .NonOverlappingStats "Sched wait time"}}</td>
	  </tr>
	{{end}}
	</table>
	`))
	type row struct {
		*GoroutineSummary
		NonOverlappingStats map[string]time.Duration
	}
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		var rows []row
		var totalExec, grandExec, maxTotal time.Duration

		for _, s := range summaries {
			grandExec += s.ExecTime // 全体の ExecTime 合計（割合計算用）
			if s.Name != name {
				continue
			} // クエリ name に一致するものだけ拾う
			rows = append(rows, row{s, s.NonOverlappingStats()})
			totalExec += s.ExecTime // グループ内の ExecTime 合計
			if s.TotalTime > maxTotal {
				maxTotal = s.TotalTime
			} // 棒グラフの基準
		}

		sort.Slice(rows, func(i, j int) bool { // 詳細表は TotalTime 降順
			return rows[i].TotalTime > rows[j].TotalTime
		})

		percent := ""
		if grandExec > 0 {
			percent = fmt.Sprintf("%.2f%%", float64(totalExec)/float64(grandExec)*100)
		}

		data := struct {
			Name            string
			N               int
			ExecTimePercent string
			MaxTotal        time.Duration
			Goroutines      []row
		}{name, len(rows), percent, maxTotal, rows}

		if err := tpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), 500)
		}
	}
}

func main() {
	summaries := fakeSummaries()
	http.HandleFunc("/goroutines", GoroutinesHandlerFunc(summaries))
	http.HandleFunc("/gorutine", GoroutinesHandler(summaries))
	log.Println("http://localhost:8080/goroutines")
	log.Fatal(http.ListenAndServe(":8080", nil))
	// Sample_API1()
	// go Sample_API2()
	// time.Sleep(1 * time.Second)

}
