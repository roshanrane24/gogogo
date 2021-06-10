package main

import(
    "fmt"
)

type hero struct{
    heroName string
}

type ua struct {
    name string
    quirk string
    height float32
    age int
    gender byte
    hero // embedding
}

func main() {
    deku := ua{
        name: "Izuku",
        quirk: "One for All",
        height:  1.66,
        age: 16,
        gender: 'M',
    }
    hero1 := hero{
        heroName: "Really Bigashlasikhfhaskl",
    }
    kachhan := ua{
        "Bakugo",
        "Explosion",
        1.72,
        16,
        'M',
        hero1,
    }

    deku.heroName = "Deku"

    fmt.Println(deku)
    fmt.Println(deku.name)
    fmt.Println(kachhan)
    fmt.Println(kachhan.height)
}
