while read line; do
	echo "Вставляем этот текст: $line"
	echo "$line" | pbcopy 		
	# sleep 5
	read -t 60 -u 1 -p 'Нажмите кнопку чтобы продолжить' key
done < test.txt
