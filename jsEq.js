class Deck {
  constructor(values) {
    this.values = values;
  }
  print() {
    this.values.forEach((card, i) => {
      console.log(i, card);
    });
  }
}

const cards = new Deck(["ace of spade", "five of diamond"]);
cards.print();
