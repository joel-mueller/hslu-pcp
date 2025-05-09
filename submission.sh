#!/bin/bash

echo "Building everything"
./package.sh

echo "Creating PDF from Bericht"
pandoc Bericht.md --pdf-engine=tectonic --metadata-file=metadata.yaml -o Bericht.pdf

echo "Creating Target Folder"
target_folder="PCP-Go-$(date -u "+%d.%m.%Y")"
mkdir "$target_folder"

echo "Copying Directories to the Target Folder"
for dir in *; do
    if [ -d "$dir" ] && [ "$dir" != "$target_folder" ]; then
        cp -R "$dir" "$target_folder/"
    fi
done

copyToTarget() {
  cp "$1" "$target_folder/"
}

echo "Moving files to the Target Directory"
copyToTarget main.go
copyToTarget go.mod
copyToTarget package.sh
copyToTarget README.md
copyToTarget Bericht.pdf

echo "Removing existing Targets"
rm *.zip

echo "Zipping the target folder"
zip -r "${target_folder}.zip" "$target_folder"

echo "Removing the unzipped target folder"
rm -rf "$target_folder"

echo "Done."