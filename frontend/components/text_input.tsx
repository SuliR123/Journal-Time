import { Dispatch, SetStateAction, useEffect, useRef, useState } from "react";
import CheckmarkIcon from "@/public/checkmark.svg"
import RestartIcon from "@/public/restart.svg"
import NoteIcon from "@/public/note.svg"
import IconButton from "./icon_button";
import ParagraphIcon from "@/public/paragraphFormat.svg"
import FocusedIcon from "@/public/focused.svg"
import Timer from "../models/timer";
import ButtonHolder from "./button_holder";
import { Format } from "@/models/format"
import SplitOverlayComponent from "./split_overlay_component";
import StatsPage from "./stats";
import SearchBar from "./search_bar";
import BoardPopup from "./board_popup";
import OverlayComponent from "./overlay_component";

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

function wpm(wordCount: number, timeSeconds: number): number {
    const wordsPerSecond = wordCount / timeSeconds
    return Math.floor(wordsPerSecond * 60)
}

function formatTime(totalSeconds: number): string {
    const hours = Math.floor(totalSeconds / 3600)
    const minutes = Math.floor(totalSeconds / 60) % 60
    const seconds = totalSeconds % 60
    return `${formatZeros(hours)}:${formatZeros(minutes)}:${formatZeros(seconds)}`
}

function formatZeros(n: number): string {
    if (n < 10 ) {
        return `0${n}`
    } else {
        return `${n}`
    }
}

interface SubmissionProps {
    title: string
    setTitle: Dispatch<SetStateAction<string>>
}

function SubmissionComponents({title, setTitle} : SubmissionProps) {
    return (
        <div className="flex flex-row w-full h-full justify-between items-center">
            <input
                id="titleBox"
                value={title}
                onChange={(e) => 
                    {
                        setTitle(e.target.value)
                    }}
                placeholder="Title..."
                className={`resize-none whitespace-nowrap overflow-x-scroll [scrollbar-width:none] [-ms-overflow-style:none] [&::-webkit-scrollbar]:hidden overflow-hidden focus:outline-none text-nowrap font-bold font-hack text-[40px] w-full h-full`} 
            />
            <SearchBar />
        </div>
    )
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
  const [title, setTitle] = useState("")

  // text input format
  const currentFormatStyle = getFormatStyle(format)

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
    <div className="relative flex items-center justify-center flex-col gap-3 h-full">
        <OverlayComponent 
            header={<SubmissionComponents title={title} setTitle={setTitle}/>}
            subheading1={<span>Date</span>} 
            subheading2={<span>Location</span>}
            show={popup} 
            body={[
                {title: "Word Count", value: `${numWords}`},
                {title: "Time Elapsed", value: formatTime(timerRef.current.getTime())},
                {title: "WPM", value: `${wpm(numWords, timerRef.current.getTime())}`}
            ]}
            onClose={() => {
                setPopup(false)
                timerRef.current.start()
            }}
            onSubmit={() => {}}
        />
        <ButtonHolder>
            <IconButton icon={<ParagraphIcon />} onClick={() => {setFormat(Format.FULL)}} displayText="full"/>
            <IconButton icon={<FocusedIcon />} onClick={() => {
                setFormat(Format.FOCUSED)}
            } displayText="focused"/> 
        </ButtonHolder>
        <span className="w-[75vw] text-[24px] text-start font-bold font-hack">{numWords}</span>
      <textarea 
        id={"textBox"}
        value={text}
        onChange={(e) => setText(e.target.value)}
        placeholder="Type something..."
        className={`${currentFormatStyle} resize-none focus:outline-none w-[75vw] text-[24px] text-start font-bold font-hack`}
      /> {/* TODO ADD CURSOR ANIMATION */}
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