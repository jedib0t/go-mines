package digital

import "strings"

var (
	numberBoxMap = map[int]string{
		-1: `
  
══
  `,
		0: `
╔═╗
║ ║
╚═╝`,
		1: `
 ╖ 
 ║ 
 ╨ `,
		2: `
╒═╗
╔═╝
╚═╛`,
		3: `
╒═╗
 ═╣
╘═╝`,
		4: `
╥ ╥
╚═╣
  ╨`,
		5: `
╔═╕
╚═╗
╘═╝`,
		6: `
╔═╕
╠═╗
╚═╝`,
		7: `
╒═╗
  ║
  ╜`,
		8: `
╔═╗
╠═╣
╚═╝`,
		9: `
╔═╗
╚═╣
╘═╝`,
	}
)

func getDigitalNumber(n int) string {
	return strings.Replace(numberBoxMap[n], "\n", "", 1)
}
