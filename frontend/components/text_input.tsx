import { useEffect, useRef, useState } from "react";
import CheckmarkIcon from "@/public/checkmark.svg"
import RestartIcon from "@/public/restart.svg"
import IconButton from "./icon_button";
import Timer from "./models/timer";

function getNumberOfWords(text: string): number {
    let words: string[] = text.split(" ");
    let count: number = 0; 
    for (let i : number = 0; i < words.length; i++) {
        let word: string = words[i];
        if(word !== " " && word !== "") {
            count++; 
        }
    } 
    return count; 
}

export default function TextInput() {
  const [text, setText] = useState("");
  const [numWords, setNumWords] = useState(0)
  const timerRef = useRef<Timer>(new Timer());
  const [startTimer, setStartTimer] = useState(false); 

  useEffect(() => {
    setNumWords(getNumberOfWords(text));
    setStartTimer(numWords > 0);
  }, [text]); 

  useEffect(() => {
    if(startTimer) {
        timerRef.current.start();
    }

    return () => {
        timerRef.current.reset(); 
    }
  }, [startTimer]); 

  return (
    <div className="flex items-center justify-center flex-col gap-3">
        <span className="w-[85vw] text-[24px] text-start font-bold font-hack">{numWords}</span>
      <textarea 
        value={text}
        onChange={(e) => setText(e.target.value)}
        placeholder="Type something..."
        className="[scrollbar-width:none] [-ms-overflow-style:none] [&::-webkit-scrollbar]:hidden overflow-hidden focus:outline-none w-[85vw] h-[8vh] text-[24px] text-start font-bold font-hack"
      />
      <div className="flex flex-row justify-between w-[10vw] h-full">
        <IconButton icon={<RestartIcon/>} onClick={() => {
            setText("");
            setStartTimer(false);
            timerRef.current.reset(); 
            }}
        />
        <IconButton icon={<CheckmarkIcon/>} onClick={() => {
            alert("DONE BUTTON TYPED FOR: " + timerRef.current.getTime());
            timerRef.current.reset(); 
            setStartTimer(false); 
            }}
        />
      </div>
      
    </div>
  );
}