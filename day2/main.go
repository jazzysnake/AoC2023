package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
  "strings"
)

type Draw struct {
  reds int
  greens int
  blues int
}

type Game struct {
  id int
  draws []Draw
}

func parseGame(input string) Game {
  trimmed := strings.TrimPrefix(input, "Game ")
  id := make([]rune,0)
  drawsBegin := 0
  for i, c := range trimmed {
    if c == ':' {
      drawsBegin=i+1
      break
    }
    id = append(id, rune(c))
  }
  idParsed, err := strconv.Atoi(string(id))
  if err != nil {
    log.Panic(err)
  }
  drawsUnparsed := strings.Split(trimmed[drawsBegin:],";") 
  draws := make([]Draw, 0)
  for _, drawStr := range drawsUnparsed {
      draws = append(draws, parseDraw(drawStr))
  }
  return Game{idParsed, draws}
}

func parseDraw(input string) Draw {
  r := 0
  g := 0
  b := 0
  colors := strings.Split(input, ",")
  for _, color := range(colors) {
    if strings.Contains(color, "red") {
      r = parseColor(color)
    }
    if strings.Contains(color, "green") {
      g = parseColor(color)
    }
    if strings.Contains(color, "blue") {
      b = parseColor(color)
    }
  }
  return Draw{r,g,b}
}

func parseColor(input string) int {
  count := make([]rune, 0)
  for _, c := range strings.TrimPrefix(input," ") {
    if c == ' ' {
      break
    }
    count = append(count, c)
  }
  res, err := strconv.Atoi(string(count))
  if err != nil {
    return 0
  }
  return res
}

func validGame(game Game, maxRed, maxGreen, maxBlue int) bool {
  for _, draw := range game.draws {
    if maxRed < draw.reds || maxBlue < draw.blues || maxGreen < draw.greens {
      return false
    }
  }
  return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
    txt := scanner.Text()
		game := parseGame(txt)
    if validGame(game, 12,13,14) {
      sum += game.id
    }
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	println(sum)
}
