#!/bin/bash

echo "Creating PDF from Bericht"
pandoc Bericht.md --pdf-engine=tectonic --metadata-file=metadata.yaml -o Bericht.pdf

echo "Creating Target Folder"
mkdir PCP-Go-$(date -u "+%d.%m.%Y %H:%M")

echo "Copying Folders to the Target Folder"
for dir in */; do
    if [ -d "$dir" ]; then
        cp -r "$dir" "$target_folder/"
    fi
done