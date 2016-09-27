all:
	go run main.go ./testdata/*.png
	benchcmp old.txt new.txt > difference.txt
	benchcmp imagemagick.txt new.txt > difference-imagemagick.txt
