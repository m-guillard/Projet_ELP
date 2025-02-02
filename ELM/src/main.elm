-- ok
module Main exposing (..)

import Browser
import Html exposing (..)
import Html.Attributes exposing (style, placeholder, value)
import Html.Events exposing (onClick, onInput)
import Svg exposing (svg)
import Svg.Attributes exposing (viewBox, width, height)
import Dessin exposing (..)  -- Utiliser Dessin pour accéder à Coord_deux_points
import Parsing exposing (..)  -- Importer Parsing pour utiliser Point
import String exposing (trim, split)
import Maybe exposing (withDefault)
import Debug exposing (log)
import Parser exposing(run)

-- MAIN

main =
    Browser.sandbox { init = init, update = update, view = view }

-- MODEL

type alias Model =
    { lines : List Dessin.Coord_deux_points  -- Utiliser Coord_deux_points : 
    , input : String  -- Ajout d'un champ pour l'entrée de texte
    }

init : Model
init =
    { lines = []
    , input = ""  -- L'entrée commence vide
    }

-- UPDATE

type Msg
    = Draw
    | UpdateInput String  -- Nouveau message pour mettre à jour l'entrée de texte

update : Msg -> Model -> Model
update msg model =
    case msg of
        Draw ->
            let
                _ = Debug.log "Input avant parsing" model.input
                parsedInstructions = run listeParser model.input
                _ = Debug.log "Résultat parsing" parsedInstructions

                instructions =
                    case parsedInstructions of
                        Ok expr ->
                            let
                                cleanedInstructions = aplatirListe expr |> List.map toFloatTuple
                            in
                            Parsing.generatePoints (aplatirListe expr)
                            -- Parsing.generatePoints cleanedInstructions
                        Err _ -> []

                _ = Debug.log "Instructions après conversion" instructions
                newPoints = instructions  -- Déjà une liste de points, pas besoin de re-appeler Parsing.generatePoints

                _ = log "Generated Points" newPoints
                newLines = createLinesFromPoints newPoints
                _ = log "Generated Lines" newLines
            in
            { model | lines = newLines }

        UpdateInput newInput ->
            { model | input = newInput }

-- Fonction pour convertir Expression en (Float, Float)
toFloatTuple : Expression -> (Float, Float)
toFloatTuple expr =
    case expr of
        Dir angle -> (toFloat angle, 0)       -- Direction (tourner) : on change l'angle
        Forward dist -> (0, toFloat dist)     -- Avancer : on change la distance parcourue
        _ -> (0, 0)                           -- Cas par défaut pour Empty ou ExprList




-- Parsing : Convertir la chaîne de caractères en liste de tuples (angle, distance)
parseInstructions : String -> List (Float, Float)
parseInstructions input =
    let
        -- Debug pour afficher l'entrée avant parsing
        _ = Debug.log "Premier Input" input

        -- Nettoyer l'entrée de ses espaces et parenthèses inutiles
        cleanedInput = 
            input
                |> String.trim
                |> String.dropLeft 1 -- Retirer le premier caractère "["
                |> String.dropRight 1 -- Retirer le dernier caractère "]"
                |> String.split "),("  -- Séparer les éléments à partir de "),("

        -- Nettoyage des parenthèses restantes
        cleanedInputNoParentheses = 
            List.map 
                (\s -> 
                    String.replace "(" "" (String.replace ")" "" s)
                ) 
                cleanedInput

        _ = Debug.log "Cleaned Input without parentheses" cleanedInputNoParentheses

    in
    cleanedInputNoParentheses
        |> List.filter (\s -> String.length s > 0) -- Filtrer les entrées vides
        |> List.map
            (\pair -> 
                case String.split "," pair of
                    [angleStr, distanceStr] -> 
                        let 
                            -- Nettoyer les espaces et extraire les valeurs
                            cleanDistanceStr = String.trim distanceStr
                            cleanAngleStr = String.trim angleStr

                            -- Debug avant la conversion
                            _ = Debug.log "Parsing angle et distance (avant conversion)" (cleanAngleStr, cleanDistanceStr)

                            -- Conversion des valeurs en float
                            angle = String.toFloat cleanAngleStr |> Maybe.withDefault 0
                            distance = String.toFloat cleanDistanceStr |> Maybe.withDefault 0

                            -- Debug après conversion
                            _ = Debug.log "Valeurs converties (après conversion)" (angle, distance)
                        in
                        (angle, distance)

                    _ -> 
                        (0, 0)  -- Si le split échoue, on renvoie une valeur par défaut
            )







-- Fonction pour créer des lignes à partir des points générés
-- Si la liste contient plusieurs points (p1 :: p2 :: rest), la fonction crée un objet de type Dessin.Coord_deux_points qui représente un segment de ligne entre les deux premiers points p1 et p2
-- Fonction pour créer des lignes à partir des points générés
createLinesFromPoints : List Parsing.Point -> List Dessin.Coord_deux_points
createLinesFromPoints points =
    case points of
        [] -> []
        [ _ ] -> []  -- Si il n'y a qu'un point, pas de ligne à tracer
        p1 :: p2 :: rest -> 
            let
                -- Créer la ligne entre p1 et p2 avec l'ordre correct des coordonnées
                firstLine = { x1 = p1.x, y1 = p1.y, x2 = p2.x, y2 = p2.y }
                _ = Debug.log "firstLine details format" ("x1: " ++ String.fromFloat p1.x ++ ", y1: " ++ String.fromFloat p1.y ++ ", x2: " ++ String.fromFloat p2.x ++ ", y2: " ++ String.fromFloat p2.y)
                -- debug pour vérifier l'ordre (x1,y1,x2,y2) car les logs affichent plutôt (x,x,y,y)
            in
            -- Ajouter le premier segment et traiter les suivants
            firstLine :: createLinesFromPoints (p2 :: rest)



-- VIEW :  l'interface utilisateur avec un champ pour saisir les instructions, un bouton pour déclencher le dessin, et un canevas pour afficher les lignes tracées
stylePage : List(Html.Attribute Msg)
stylePage =
    [ style "box-sizing" "border-box"
    , style "width" "100%"
    , style "max-width" "500px"
    , style "margin" "auto"
    , style "padding" "1rem"
    , style "display" "grid"
    , style "grid-gap" "30px"
    ]

view : Model -> Html Msg
view model =
    div 
        [ style "font-family" "-apple-system, BlinkMacSystemFont, Segoe UI, Roboto, Open Sans, Ubuntu, Fira Sans, Helvetica Neue, sans-serif"
        , style "margin" "0"
        , style "min-height" "100vh"
        ]
        [ div
            stylePage
            [ text "Type in your code below:" ]
        , div
            stylePage
            [ input 
                [ placeholder "example: [Repeat 360 [Forward 1, Left 1]]"
                , value model.input
                , onInput UpdateInput
                ]
                []  -- Supprime l'erreur "Too many args"
            , button [ onClick Draw ] [ text "Draw" ]
            , svg
                [ style "width" "100%"
                , style "max-width" "500px"
                , style "border" "1px solid #D1C4E9"
                , viewBox "0 0 500 500"
                , width "500"
                , height "500"
                ]
                (Dessin.display (transformLinesToPoints model.lines))
            ]
        ]

