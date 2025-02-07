// 99 Bottles Of Beer, Golang Buffered Channels Version
// 
// Mehmet Gürbüz GÜVEN
// https://github.com/gurbuzguven
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
package main

import (
    "fmt"
    "strconv"
)

func fillBottles(bottlesOfBeer int, bottles chan<- int){
    // Fill up the bottles' buffer
    for bottlesOfBeer >= 0 {
        bottles <- bottlesOfBeer
        bottlesOfBeer--
    }
    
    // Close the bottles channel when we're done
    close(bottles)

}

func prepareLyrics(bottles <-chan int, lyrics chan<- string) {

    // Receive from bottles' buffer
    for bottle := range bottles {
        // Run lyricsRefrain and send to lyrics channel
        lyrics <- lyricsRefrain(bottle)
    }
    
    // Close the lyrics channel when we're done
    close(lyrics)
}



func lyricsRefrain(bottle int) string { 
   
    switch bottle {
    
        case 1:
            return "1 bottle of beer on the wall, 1 bottle of beer.\nTake one down and pass it around, no more bottles of beer on the wall.\n\n"
            
        case 0:
            return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
            
        default:
            return strconv.Itoa(bottle) + " bottles of beer on the wall, " + strconv.Itoa(bottle)  + " bottles of beer.\nTake one down and pass it around, " + strconv.Itoa(bottle - 1) + " bottles of beer on the wall.\n\n"
            
    }
}

func printLyrics(lyrics <-chan string) {
    // Receive and print from lyrics' buffer
    for lyric := range lyrics {
        fmt.Print(lyric)
    }
}


func main() {
   
    bottlesOfBeer := 99
    
    // Create channels with buffers
    bottles, lyrics := make(chan int, bottlesOfBeer), make(chan string, bottlesOfBeer)
    
    go fillBottles(bottlesOfBeer, bottles)
    go prepareLyrics(bottles, lyrics)
    printLyrics(lyrics)
}
