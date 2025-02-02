module Parsing exposing (..)

import Parser exposing (..)


type Expression
    = Dir Int
    | Forward Int
    | ExprList (List Expression)
    | Empty


dupliquer : Int -> Expression -> List Expression
dupliquer n repet =
    if n <= 0 then
        []
    else
        case repet of
            ExprList lst -> lst ++ dupliquer (n - 1) (ExprList lst)
            _ -> repet :: dupliquer (n - 1) repet

-- Evite d'avoir des listes imbriquées
aplatirListe : Expression -> List Expression
aplatirListe expr =
    case expr of
        ExprList lst -> List.concatMap aplatirListe lst
        Empty -> []
        _ -> [expr]

listeParser : Parser Expression
listeParser =
    succeed (\expr rest ->
        case rest of
            ExprList lst -> ExprList (expr :: lst)
            Empty -> ExprList [expr]
            _ -> ExprList [expr, rest] 
    )
        |. symbol "["
        |. spaces
        |= expressionParser
        |= oneOf [ lazy (\_ -> restParser), succeed Empty ]
        |. spaces
        |. symbol "]"


restParser : Parser Expression
restParser =
    succeed (\expr rest ->
        case rest of
            ExprList lst -> ExprList (expr :: lst)  -- Fusionne l'élément avec la liste existante
            Empty -> ExprList [expr]  -- Cas où il n'y a pas de suite
            _ -> ExprList [expr, rest]  -- Si `rest` est un élément simple, crée une liste avec lui
    )
        |. symbol ","
        |. spaces
        |= expressionParser
        |= oneOf [ lazy (\_ -> restParser), succeed Empty ]

repeatParser : Parser Expression
repeatParser =
    succeed (\n expr -> ExprList (dupliquer n expr))
        |. symbol "Repeat"
        |. spaces
        |= int
        |. spaces
        |= lazy (\_ -> listeParser)

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
    token "Forward"
    |> andThen (\_ ->
        spaces
            |> andThen (\_ ->
                int
                    |> andThen (\n ->
                        succeed (Forward n)
                    )
            )
    )

expressionParser : Parser Expression
expressionParser =
    oneOf [ repeatParser, rightParser, leftParser, forwardParser ]

test : String
test =
    let
        inputString = "[Repeat 8 [Left 45, Repeat 6 [Repeat 90 [Forward 1, Left 2], Left 90]]]"
        result =
            case run listeParser inputString of -- On exécute le parser
                Ok expr -> Debug.toString (aplatirListe expr) -- On transforme la liste
                Err _ -> "Erreur parsing"
    in
    result

    