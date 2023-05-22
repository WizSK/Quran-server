#!/bin/bash

# https://api.qurancdn.com/api/qdc/verses/by_chapter/2?words=true&translation_fields=resource_name%2Clanguage_id&per_page=5&fields=text_uthmani%2Cchapter_id%2Chizb_number%2Ctext_imlaei_simple&translations=131%2C180%2C20&reciter=7&word_translation_language=en&page=1&word_fields=verse_key%2Cverse_id%2Cpage_number%2Clocation%2Ctext_uthmani%2Ccode_v1%2Cqpc_uthmani_hafs&mushaf=2
# https://api.qurancdn.com/api/qdc/verses/by_chapter/${index}?words=true&word_translation_language=en&page=1


for i in $(seq 1 114);
do
    # pages=$(sed -n ${i}p pages.txt)
    # strt=$(echo $pages | cut -d " " -f 2)
    # finish=$(echo $pages | cut -d " " -f 3)

    mkdir -p "$i"
    echo $i

    idx=1
    file="${i}/${idx}.json"
    curl -s "https://api.qurancdn.com/api/qdc/verses/by_chapter/${i}?words=true&word_translation_language=bn" > "$file"
    next=$(cat "$file" | jq ".pagination.next_page")
    echo "$(cat ${file} |  jq '.pagination.total_pages' )" > "${i}/page_count.txt"
    
    while [ "$next" != "null" ]
    do
        idx=$((idx+1))
        file="${i}/${idx}.json"

        curl -s "https://api.qurancdn.com/api/qdc/verses/by_chapter/${i}?words=true&word_translation_language=bn&page=${idx}" > "${i}/${idx}.json"
        next=$(cat "$file" | jq ".pagination.next_page")
        echo "    $idx"
    done
    echo "$next"
done
