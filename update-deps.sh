for folder in $(find . -mindepth 1 -maxdepth 1 -type d \( -name "*" \) );
do
  cd "$folder" || exit

  echo "Folder: $folder"
  #go get -u -v golang.org/x/net
  #go get -u -v golang.org/x/crypto
  go get -u -v 

  cd -
done
