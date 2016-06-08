for lettera in Q W E R T Y U I O P A S D F G H J K L Z X C V B N M
do
export max="$(wget "http://www.cognomix.it/origine-cognomi-italiani/$lettera" 2>/dev/null -O- | grep -Po "href=\"http://www.cognomix.it/origine-cognomi-italiani/$lettera/\K[0-9]+" | tail -n 1 )"
for numero in `seq $max 2>/dev/null || echo 1`
do
	echo Request: http://www.cognomix.it/origine-cognomi-italiani/$lettera/$numero 1>&2
wget "http://www.cognomix.it/origine-cognomi-italiani/$lettera/$numero" 2>/dev/null -O- |
grep -Po 'php" title="\K[^-]+'
done
done > cognomix.txt
