-- checker les appelations des fichiers lors du merge de nos deux travaux
module Parsing_Alice exposing (..)

import Basics exposing (pi, cos, sin)

-- Définir un alias pour un point (x, y)
type alias Point =
    { x : Float, y : Float }

-- Fonction pour calculer la nouvelle position après un déplacement (angle, distance) : [(90, 20), (140, 10), (-55, 20)] -> lsite de points [{x = 0, y = 0}, {x = 0, y = 20}, {x = 7.66, y = 25.88}, {x = 2.91, y = 42.91}]
-- Mise à jour de `generatePoints` pour ignorer les déplacements nuls
generatePoints : List (Float, Float) -> List Point
generatePoints instructions =
    let
        (points, _) =
            List.foldl
                (\(angle, distance) (acc, prevPoint) ->
                    let
                        _ = Debug.log "Point actuel" prevPoint
                        _ = Debug.log "Processing Instruction dans generatePoints" (angle, distance)

                        newPoint =
                            if distance == 0 then
                                prevPoint
                            else
                                calculateNewPosition prevPoint (angle, distance)

                        _ = Debug.log "Nouveau point calculé" newPoint
                    in
                    (acc ++ [newPoint], newPoint)
                )
                ([], { x = 250, y = 250 })
                instructions
    in
    { x = 250, y = 250 } :: points  -- On ajoute manuellement le point d'origine en premier



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






