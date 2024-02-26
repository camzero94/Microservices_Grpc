#!/bin/bash

start_server(){
    # Build Phase 
    echo "================== Building...... ====================="
    echo "$1"
    # Compile svc1 and svc2
    cd "$1"/ && make build && cd ..
    cd "$2"/ && make build && cd ..

    # Build Running Phase
    echo "================== Running...... ====================="
    cd "$1"/ && make run &
    pid_svc1=$!
    cd "$2"/ && make run &
    pid_svc2=$!

    # Wait it finish background process
    wait $pid_svc1
    wait $pid_svc2

    echo "All servers have closed."
}

clean_dir(){
    echo "The folder is $1, $2 " 
    rm -rf "$1/bin" "$2/bin"
}

show_menu(){
    echo "This are the dir are going to be run "$1" and "$2". Are you sure you want to run it." 
    read user_answer
    options=("Run the servers" "Delete the dir binaries" "Quit")

    select choice in "${options[@]}";do
        case $REPLY in
            1)
                echo "You select: $choice"
                start_server "$1" "$2"
                ;;
            2)
                echo "You select: $choice"
                clean_dir "$1" "$2"
                ;;
            3)
                echo "You select Quit the program."
                exit 0
                ;;
            *)
                echo "Invalid input. Please choose number beetween 1 and ${#options[@]}"
                exit 

        esac
    done
}



show_menu "$1" "$2"     

