# Visualization stuff

go run main/main.go
sfdp -x -Goverlap=scale -Tpng graph.txt > graph.png
convert graph.png -resize 4000x4000 graph-resize.png

## live preview

```
watchexec -e go,kage -r 'go run part12.go && neato -Tpng graph.txt > graph.png && /Users/aros/.iterm2/imgcat graph.png'
```

```
