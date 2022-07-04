# Discovery

Challenge:

> Create a REST API to handle the deck and cards to be used in any card game.

There are a few important points in this sentence:

1. _REST API_

    The first thing I think about this topic is: REST equals CRUD for web. Next I think of RESTFul Routing and Entity Modeling.

2. _Deck and cards_

    Thinking about entities: decks and cards. Deck is a collection of cards. All the deck entity has is an ID, a set of cards, and some transient data. On the other hand, the cards have names (values), suits and codes.

3. _Any card game_

    I assume that we are talking about card games played with the standard 52-card deck of French playing cards.

## Thinking about the non-functional requirements

This project aims to be:

- easy to extend
- easy to modify
- easy to understand to others

_How can it be easy to extend?_ I think developing abstraction layers is a good way out when semantically combined with the project structure. Like `package model` or `package repository`.

_How can it be easy to modify?_ To me, this can be thought of as the ease by which I can implement a new endpoint or a query param. There are a few approaches to achieving this, the one I'm going to use is to have a clear and simple flow. As in: `model -> service -> api`, a change to the model would cause changes to the service and probably the API. I think for a small project this approach is OK, but it could be better.

_How can it be easy to understand to others?_ Unified flow, good practices and patterns.

## Thinking about the design

A deck can be represented as a document or as a table and, in order to do that, I think I can use Postgres as SSoT when we talk about the representation of the deck. _Why?_ Fast to query and reliability.

Redis seems like a good alternative for handling hot data. _How?_ I'll take the table in Postgres as a mirror of data that I want (or not) to put in Redis as in `key = deck_id` and `value = model.Deck`.
