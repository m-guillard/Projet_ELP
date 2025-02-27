-- OK - elm install elm/parser
module Parsing exposing (..)
import Parser exposing (..)
import Basics exposing (pi, cos, sin)


-- Définir un alias pour un point (x, y)
type alias Point =
    { x : Float, y : Float }

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

-- Fonction pour calculer la nouvelle position après un déplacement (angle, distance) : [(90, 20), (140, 10), (-55, 20)] -> lsite de points [{x = 0, y = 0}, {x = 0, y = 20}, {x = 7.66, y = 25.88}, {x = 2.91, y = 42.91}]
-- Mise à jour de `generatePoints` pour ignorer les déplacements nuls


generatePoints : List Expression -> List Point
generatePoints instructions =
    let
        (points, _, _) =
            List.foldl
                (\expr (acc, prevPoint, currentAngle) ->
                    case expr of
                        Forward dist ->
                            let
                                newPoint = calculateNewPosition prevPoint (currentAngle, toFloat dist)
                                _ = Debug.log "Nouveau point calculé" newPoint
                            in
                            (acc ++ [newPoint], newPoint, currentAngle)

                        Dir angle ->
                            -- Mise à jour de l'angle seulement, sans ajouter de point
                            (acc, prevPoint, currentAngle + toFloat angle)

                        _ ->
                            (acc, prevPoint, currentAngle)
                )
                ([], { x = 250, y = 250 }, 0) -- Angle initial à 0°
                instructions
    in
    { x = 250, y = 250 } :: points  -- Ajouter le point de départ


calculateNewPosition : Point -> (Float, Float) -> Point
calculateNewPosition { x, y } (angle, distance) =
    let
        -- Convertir l'angle en radians
        radians = angle * pi / 180

        -- Calculer les nouvelles coordonnées
        newX = x + (distance * cos radians)
        newY = y + (distance * sin radians)
    in
    { x = newX, y = newY }






