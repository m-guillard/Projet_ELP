module ParsingNew exposing (..)

import Parser exposing (succeed, token, int, spaces, lazy, run)
import Array exposing (empty)
import Maybe exposing (andThen)


type Expression
    = Repeat Int (List Expression)
    | Right Int
    | Left Int
    | Forward Int

-- à faire
parseRepeat : Parser.Parser Expression
parseRepeat =
    token "Repeat"
        |> Parser.andThen (\_ ->
            spaces
                |> andThen (\_ ->
                    int
                    |> Parser.andThen (\n ->
                        spaces
                        |> Parser.andThen (\_ ->
                            token "["
                            |> Parser.andThen (\_ ->
                                parseListe
                                |> Parser.andThen (\exprs ->
                                    token "]"
                                        |> succeed (Fois n exprs)
                                )
                            )
                        )
                    )
                )
        )

parseRight : Parser.Parser Expression
parseRight =
    token "Right"
        |> Parser.andThen (\_ ->
            spaces
                |> andThen (\_ ->
                    int
                        |> Parser.andThen (\n -> 
                            succeed (Right n)
                        )
                )
        )

parseLeft : Parser.Parser Expression
parseRight =
    token "Left"
        |> Parser.andThen (\_ ->
            spaces
                |> andThen (\_ ->
                    int
                        |> Parser.andThen (\n -> 
                            succeed (Left n)
                        )
                )
        )

parseForward : Parser.Parser Expression
parseForward =
    token "Forward"
        |> Parser.andThen (\_ ->
            spaces
                |> andThen (\_ ->
                    int
                        |> Parser.andThen (\n -> 
                            succeed (Forward n)
                        )
                )
        )

parseListe : Parser.Parser (List Expression)
parseListe =
    Parser.oneOf
        [ token "]"
            |> succeed []
        , parseExpression
            |> andThen (\expr ->
                spaces
                    |> andThen (\_ ->
                        optional (token ",")
                            |> andThen (\_ ->
                                spaces
                                    |> andThen (\_ ->
                                        parseListe
                                            |> andThen (\rest ->
                                                succeed (expr :: rest)
                                            )

                                    )
                            )
                    )
            )
        ]

-- On prend un truc du type [Repeat 3 [Direction nb, Forward Nb]]

-- On détecte Repeat
repeatParser : Parser.Parser String
repeatParser =
    Parser.oneOf
        token "Repeat"

duplication : String -> String
duplication input =
    repeatParser input


-- On récupère 3 qui se situe après repeat
-- On duplique ce qu'il y a dans les crochets juste après

-- Devient [Direction nb, Forward Nb], [Direction nb, Forward Nb], [Direction nb, Forward Nb]

-- Devient [-+nb, Nb],[-+nb, Nb],[-+nb, Nb]