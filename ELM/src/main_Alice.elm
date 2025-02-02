-- checker les appelations des fichiers lors du merge de nos deux travaux
module main_Alice exposing (..)

import Browser
import Html exposing (..)
import Html.Attributes exposing (style, placeholder, value)
import Html.Events exposing (onClick, onInput)
import Svg exposing (svg)
import Svg.Attributes exposing (viewBox, width, height)
import Dessin_Alice exposing (..)  -- Utiliser Dessin2 pour accéder à Coord_deux_points
import Parsing_Alice exposing (..)  -- Importer Parsing pour utiliser Point
import String exposing (trim, split)
import Maybe exposing (withDefault)
import Debug exposing (log)

-- MAIN

main =
    Browser.sandbox { init = init, update = update, view = view }

-- MODEL

type alias Model =
    { lines : List Dessin2.Coord_deux_points  -- Utiliser Coord_deux_points : 
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
update msg model = -- prend en entrée un message (msg) et un modèle (model), et retourne un nouveau modèle mis à jour en fonction du type de message reçu
    case msg of
        Draw ->
            let
                -- Parser les instructions: conversion du format texte donné par l'utilisateur [(90, 20), (140, 10), (-55, 20)], en vrais tuples (angle, dist)
                instructions = parseInstructions model.input -- fonction de Main
                _ = Debug.log "Instructions après parsing" instructions
                _ = log "Instructions parsed" instructions

                -- Générer les points : calculer la nouvelle position après un déplacement (angle, distance) : par exemple, [(90, 20), (140, 10), (-55, 20)] -> lsite de points [{x = 0, y = 0}, {x = 0, y = 20}, {x = 7.66, y = 25.88}, {x = 2.91, y = 42.91}]
                _ = Debug.log "Instructions parsed, juste avant d'appeler generatePoints dans le main" instructions
                newPoints = Parsing.generatePoints instructions
                _ = log "Generated Points" newPoints

                -- Générer les lignes à partir des points: Entrée : liste de points [ {x = 0, y = 0}, {x = 10, y = 10}, {x = 20, y = 20} ] ; Sortie : [ {x1 = 0, y1 = 0, x2 = 10, y2 = 10}, {x1 = 10, y1 = 10, x2 = 20, y2 = 20} ]
                newLines = createLinesFromPoints newPoints -- fonction de Main
                _ = log "Generated Lines" newLines


            in
            { model | lines = newLines }  -- Mettre à jour les lignes avec les nouvelles lignes générées : Le modèle contient une liste de lignes (lines) qu'on remplace par newLines

        UpdateInput newInput -> -- Màj la valeur du champ input dans le modèle avec la nouvelle valeur newInput rentrée par l'user, l'événement onInput UpdateInput est déclenché.
            { model | input = newInput }

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
-- Si la liste contient plusieurs points (p1 :: p2 :: rest), la fonction crée un objet de type Dessin2.Coord_deux_points qui représente un segment de ligne entre les deux premiers points p1 et p2
-- Fonction pour créer des lignes à partir des points générés
createLinesFromPoints : List Parsing.Point -> List Dessin2.Coord_deux_points
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

view : Model -> Html Msg
view model =
    div []
        [ input
            [ placeholder "Exemple: [(90, 20),(140, 10), (-55, 20)]"
            , value model.input
            , onInput UpdateInput
            ]
            []
        , button [ onClick Draw ] [ text "Draw" ]
        , svg
            [ style "width" "100%"
            , style "max-width" "500px"
            , style "border" "1px solid #D1C4E9"
            , viewBox "0 0 500 500"
            , width "500"
            , height "500"
            ]
            (Dessin2.display (transformLinesToPoints model.lines))  -- Convertir les lignes en points  -- Passer model.lines (de type Coord_deux_points) directement ici, et non pas du Point
            -- (Dessin2.display (transformLinesToPoints model.lines))
        ]

