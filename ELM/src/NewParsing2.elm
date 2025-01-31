module NewParsing2 exposing (..)

import Parser exposing (Parser, oneOf, andThen, succeed, token, int, spaces, lazy, run)

type Expression
    = Dir Int
    | Forward Int
    | Nothing

dupliquer : Int -> String -> String -> String
dupliquer n repet txt =
    case n of
        0 -> txt
        _ -> dupliquer (n - 1) repet (txt ++ repet)

arrangement : Expression -> Expression -> String
arrangement exp1 exp2 =
    case (exp1,exp2) of
        (Dir n1, Forward n2)  -> (String.fromInt n1) ++ " " ++ (String.fromInt n2)
        (Forward n1, Dir n2)  -> (String.fromInt n2) ++ " " ++ (String.fromInt n1)
        (Nothing, Forward n2)  -> (String.fromInt 0) ++ " " ++ (String.fromInt n2)
        (Dir n1, Nothing)  -> (String.fromInt n1) ++ " " ++ (String.fromInt 0)
        (Nothing, Dir n2)  -> (String.fromInt n2) ++ " " ++ (String.fromInt 0)
        (Forward n1, Nothing)  -> (String.fromInt 0) ++ " " ++ (String.fromInt n1)
        (Forward n1, Forward n2) -> (String.fromInt 0) ++ " " ++ (String.fromInt n1) ++ " " ++ (String.fromInt 0) ++ " " ++ (String.fromInt n2)
        (Dir n1, Dir n2) -> (String.fromInt n1) ++ " " ++ (String.fromInt 0) ++ " " ++ (String.fromInt n2) ++ " " ++ (String.fromInt 0)
        (Nothing, Nothing) ->  (String.fromInt 0) ++ " " ++ (String.fromInt 0)

listeParser : Parser String
listeParser =
    token "["
        |> andThen (\_ ->
            spaces
                |> andThen (\_ ->
                    expressionParser
                        |> andThen (\expr ->
                            spaces
                                |> andThen (\_ ->
                                    token "]"
                                        |> andThen (\_ ->
                                            succeed (expr)
                                        )
                                )
                        )
                )
        )

repeatParser : Parser String
repeatParser = 
    token "Repeat"
        |> andThen (\_ ->
            spaces
                |> andThen (\n ->
                    int
                        |> andThen (\_ ->
                            spaces
                                |> andThen (\expr ->
                                    listeParser
                                    |> andThen (\_ ->
                                        succeed (dupliquer n expr)
                                    )
                                )
                        )
                )
        )

directionParser : Parser String
directionParser =
    firstElementParser
        |> andThen (\expr1 ->
            spaces
                |> andThen (\_ ->
                    token ","
                        |> andThen (\_ ->
                            spaces
                                |> andThen (\_ ->
                                    secondElementParser
                                        |> andThen (\expr2 ->
                                            spaces
                                                |> andThen (\_ ->
                                                    succeed (arrangement expr1 expr2)
                                                )
                                        )
                                )
                        )
                )
        )

secondElementParser : Parser Expression
secondElementParser =
    oneOf
        [ rightParser
        , leftParser
        , forwardParser
        , token "Repeat"
            |> andThen (\_ ->
                succeed (Nothing)
            )
        ]

rightParser : Parser Expression
rightParser =
    token "Right"
    |> andThen (\_ ->
        spaces
            |> andThen (\_ ->
                int
                    |> andThen (\n ->
                        succeed (Dir n)
                    )
            )
    )

leftParser : Parser Expression
leftParser =
    token "Left"
    |> andThen (\_ ->
        spaces
            |> andThen (\_ ->
                int
                    |> andThen (\n ->
                        succeed (Dir -n)
                    )
            )
    )

forwardParser : Parser Expression
forwardParser =
    token "Right"
    |> andThen (\_ ->
        spaces
            |> andThen (\_ ->
                int
                    |> andThen (\n ->
                        succeed (Forward n)
                    )
            )
    )


firstElementParser : Parser Expression
firstElementParser =
    oneOf [ rightParser, leftParser, forwardParser]

expressionParser : Parser String
expressionParser =
    oneOf [ repeatParser, directionParser ]