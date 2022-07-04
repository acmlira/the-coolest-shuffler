# The Coolest Shuffler

![The standard 52-card deck of French playing cards illustration](assets/the-coolest-shuffler.png)

## Discovery

Challenge:

> Create a REST API to handle the deck and cards to be used in any card game

There are a few important points in this sentence:

1. _REST API_

    The first thing I think about this topic is: REST equals CRUD for web. Next I think of RESTFul Routing and Entity Modeling.

2. _Deck and cards_

    Thinking about entities: decks and cards. Deck is a collection of cards. All the deck entity has is an ID, a set of cards, and some transient data. On the other hand, the cards have names (values), suits and codes.

3. _Any card game_

    I assume that we are talking about card games played with the standard 52-card deck of French playing cards.
