find . '(' -name '*.sh' ')' -print | sed 's/\(.*\)\///g' | sed 's/\.sh//g'for d in $(pwd)
do
        cd "$d" || exit
        INTERVIEWNUMBER=699607
        echo "$INTERVIEWNUMBER"
        cat interviews/interview-"$INTERVIEWNUMBER"
        grep -A 4 L337 vehicles | grep -A 3 -B 1 Honda | grep -A 2 -B 2 Blue | grep -B 4 "Height: 6"
        cat memberships/AAA memberships/Delta_SkyMiles memberships/Museum_of_Bash_History memberships/Terminal_City_Library| grep "$MAIN_SUSPECT" | wc -l
done
#command used to find all the interview files that contains useful informations about the suspect 
#grep -Eo '[0-9\]+' *
