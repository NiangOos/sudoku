/*

Droit d'auteur réservé : magniang , jfreira, Halune
Copyright tout droit réservés

*/

package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

// ///////////////////////////////////////////////
//
//	ÉTAPE 1                               //
//
// //////////////////////////////////////////////
// Remplissage des parametres passés dans un tableau
func RemplissageTable(table [9][9]rune, args []string) [9][9]rune { // return un tableau de rune
	for i := range args {
		// cahque valeur de i e.g ".0.967.12"
		for j := range args[i] {
			table[i][j] = rune(args[i][j])
		}
	}
	return table
}

// ///////////////////////////////////////////////
//
//	ÉTAPE 2                               //
//
// //////////////////////////////////////////////
// pour verifier si les paramètres rentrées sont ok
func VerifiParams(args []string) bool {
	if len(args) != 9 { // longueur des arguments non egam a 9
		fmt.Println("Error") // Si les donnees rentrées sont pas égales à 9 c'est faux donc error
		return false
	}

	for i := 0; i < len(args); i++ {
		if len(args[i]) != 9 {
			fmt.Println("Error") //  Si dans chaque donnees rentrée ont n'a pas 9 éléments c'est faux donc error
			return false
		}
	}
	for i := 0; i < len(args); i++ {
		for _, value := range args[i] {
			if value == 47 || value == 48 { // si la valeur == / ou 0
				fmt.Println("Error") // Entree incorecct
				return false
			} else if value < 46 || value > 57 { // si la valeur n'est pas conprise entre . et 9;
				fmt.Println("Error") // Entree incorecct
				return false
			}
		}
	}
	return true // Si non tout est vrai
}

// ///////////////////////////////////////////////
//
//	ÉTAPE 3                               //
//
// //////////////////////////////////////////////
// Comptage des cellules vides en l'occurence des .
// On passe l'adresse du tableau il verifie puis renvoie un bool
// Si on place un * devant un nom de pointeur on obtient la valeur de la variable stockeé à l'adresse indiqué par le pointeur
func CellVide(table *[9][9]rune) bool {
	for i := 0; i < 9; i++ { // parcours la ligne
		for j := 0; j < 9; j++ { // parcours la colonne
			if table[i][j] == '.' { // si a une ligne et colonne il ya une cellule vide cest ok
				return true
			}
		}
	}
	return false // sinon faux
}

// ///////////////////////////////////////////////
//
//	ÉTAPE   4                             //
//
// //////////////////////////////////////////////
// Pour verifier si une cellule est valide
func EstValide(table *[9][9]rune, col int, ligne int, val rune) bool {
	// verification d'un double int
	// si la valeur est absente sur une colonne
	for i := 0; i < 9; i++ {
		if val == table[i][col] {
			return false
		}
	}

	// si une valeur est absente sur une ligne
	for j := 0; j < 9; j++ {
		if val == table[ligne][j] {
			return false
		}
	}

	// verification d'un carre
	a := col / 3 // on divise par 3 pour la ligne et pour la colonne pour avoir un carre
	b := ligne / 3

	for k := 3 * a; k < 3*(a+1); k++ {
		for l := 3 * b; l < 3*(b+1); l++ {
			if val == table[l][k] {
				return false
			}
		}
	}
	return true
}

/////////////////////////////////////////////////
//      ÉTAPE   5                             //
////////////////////////////////////////////////
// On utilise ici le bactracking càd on revient sur nos traces pour pouvoir resoudre un probleme

func EstResolu(table *[9][9]rune) bool { // on passe l'adresse du tableau ça renvoie un boolen
	if !CellVide(table) {
		return true // si la cellule non vide donc vrai
	}
	for ligne := 0; ligne < 9; ligne++ { // on incremente ligne
		for col := 0; col < 9; col++ { // on incremente colonne
			if table[ligne][col] == '.' { // si a une ligne et a une colonne donnée on a point càd vide
				for val := '1'; val <= '9'; val++ {
					if EstValide(table, col, ligne, val) { // si c'est valide (la cellule)
						table[ligne][col] = val // dans ce cas on met la valeur
						if EstResolu(table) {
							return true // et on met resolu à vrai comme ça la valeur ne sera pas change
						}
					}
					table[ligne][col] = '.' // si non valide on laisse à vide
				}
				return false // et on return false
			}
		}
	}
	return false
}

func main() {
	arguments := os.Args[1:] // On recupere les arguments apres le main de 1 au dernier

	if VerifiParams(arguments) == true { // verification if all ok
		table := [9][9]rune{}                      // on cree le tableau
		table = RemplissageTable(table, arguments) // on le remplie grace a notre fonction

		if EstResolu(&table) == true { // si cell résolu
			for ligne := 0; ligne < 9; ligne++ {
				for col := 0; col < 9; col++ {
					if col != 8 {
						z01.PrintRune(rune(table[ligne][col])) // on print la cellule et la colone resolue et ainsi de suite
						z01.PrintRune(32)                      // espace
					} else {
						z01.PrintRune(rune(table[ligne][col]))
					}
				}
				z01.PrintRune(10) // nouvelle ligne
			}
			z01.PrintRune(10)
		} else {
			fmt.Println("Error") // VerifiParams est fausse
		}
	}
}
