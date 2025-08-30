package chanjs

import "strings"

func matchPattern(pattern, name string) bool {
    patParts := strings.Split(pattern, ".")
    nameParts := strings.Split(name, ".")
    // Implementa una versión simple del wildcards (‘*’ = un segmento, ‘#’ = varios segmentos)
    // Por simplicidad:
    // '*' coincide con cualquier parte excepto vacío
    // '#' coincide con todo lo que queda
    for i, part := range patParts {
        if part == "#" {
            return true
        }
        if i >= len(nameParts) {
            return false
        }
        if part != "*" && part != nameParts[i] {
            return false
        }
    }
    return len(patParts) == len(nameParts)
}
