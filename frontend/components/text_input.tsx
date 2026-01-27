import { useEffect, useRef, useState } from "react";
import CheckmarkIcon from "@/public/checkmark.svg"
import RestartIcon from "@/public/restart.svg"
import NoteIcon from "@/public/note.svg"
import IconButton from "./icon_button";
import Timer from "../models/timer";
import ButtonHolder from "./button_holder";
import { Format } from "@/models/format"
import OverlayComponent from "./overlay_component";
import StatsPage from "./stats";
import SearchBar from "./search_bar";

function getNumberOfWords(text: string): number {
    let words: string[] = text.split(/\s+/);
    let count: number = 0; 
    for (let i : number = 0; i < words.length; i++) {
        let word: string = words[i];
        if(word !== " " && word !== "") {
            count++; 
        }
    } 
    return count; 
}

function getFormatStyle(format : Format): string {
    switch (format) {
        case Format.FOCUSED:
            return "[scrollbar-width:none] [-ms-overflow-style:none] [&::-webkit-scrollbar]:hidden overflow-hidden h-[8vh]";
        case Format.FULL:
            return "h-[45vh]";
        default:
            return ""
    }
}

export default function TextInput() {
  
  // text states  
  const [text, setText] = useState("");
  const [numWords, setNumWords] = useState(0)
  const [format, setFormat] = useState(Format.FOCUSED)
  
  // timer states 
  const timerRef = useRef<Timer>(new Timer());
  const [startTimer, setStartTimer] = useState(false); 

  // popup 
  const [popup, setPopup] = useState(false)

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

  // TODO: When going from full to focused if the user scrolled, it might refocus weird
  return (
    <div className="relative flex items-center justify-center flex-col gap-3">
        <OverlayComponent show={popup} 
            left={<StatsPage seconds={timerRef.current.currentTime} wordCount={numWords}/>} 
            right={<SearchBar />} 
            toggle={setPopup}
            onSubmit={() => 
                {
                    timerRef.current.reset()
                    setStartTimer(false); 
                    setText("")
                }
            }
            onClose={() => timerRef.current.start()}
        />
        <ButtonHolder>
            <IconButton icon={<NoteIcon />} onClick={() => {setFormat(Format.FULL)}} displayText="full"/>
            <IconButton icon={<NoteIcon />} onClick={() => {
                setFormat(Format.FOCUSED)}
            } displayText="focused"/> 
        </ButtonHolder>
        <span className="w-[85vw] text-[24px] text-start font-bold font-hack">{numWords}</span>
      <textarea 
        id={"textBox"}
        value={text}
        onChange={(e) => setText(e.target.value)}
        placeholder="Type something..."
        className={`${getFormatStyle(format)} focus:outline-none w-[85vw] text-[24px] text-start font-bold font-hack`}
      />
      <ButtonHolder transparent={true}>
        <IconButton icon={<RestartIcon/>} onClick={() => {
            setText("");
            setStartTimer(false);
            timerRef.current.reset(); 
            }}
        />
        <IconButton icon={<CheckmarkIcon/>} onClick={() => {
            if(numWords == 0) {
                return
            }
            setPopup(true)
            timerRef.current.pause()
            }}
        />
      </ButtonHolder>
      
    </div>
  );
}