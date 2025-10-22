use std::collections::HashSet;

// HashSet adalah kumpulan unik — tidak bisa punya elemen duplikat.

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    let normalized = normalize(word);

    possible_anagrams
        .iter() // Mengambil iterator dari slice possible_anagrams. .iter() memberikan &&str (reference ke &str),
        .filter(|&&candidate| { // filter digunakan untuk memilih elemen-elemen yang memenuhi kondisi tertentu
            //kita tulis |&&candidate| untuk langsung men-dereference dua kali.
            let candidate_norm = normalize(candidate); 
            candidate_norm == normalized &&
            word.to_lowercase() != candidate.to_lowercase()
        }) //menghasilkan iterator berisi &&str.
        .cloned() // mengubah tiap &&str menjadi &str.
        .collect()
}

fn normalize(s: &str) -> String {
    let mut chars: Vec<char> = s.to_lowercase().chars().collect();

    chars.sort_unstable(); // metode untuk mengurutkan elemen di tempat, dia langsung mengubah urutan elemen di vektor tersebut tanpa membuat salinan baru
    chars.into_iter().collect() // Mengubah hasil iterator menjadi sebuah HashSet.
}