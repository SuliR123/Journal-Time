import SearchIcon from "@/public/search.svg"
import { useState } from "react"

interface SearchBarProps {
    background?: boolean
}

export default function SearchBar({ background = false } : SearchBarProps) {
    const [searchText, setSearch] = useState("")
    const backgroundColor = background ? "bg-background-color backdrop-brightness-125" : ""
    const brightness = background ? "backdrop-brightness-125" : ""

    return (
        <div className={`h-8 ${backgroundColor} rounded-sm`}>
            <div className={`flex flex-row justify-center items-center fill-icon-color hover:fill-text-color gap-3 h-8 ${brightness}`}>
                <SearchIcon className={`w-6 h-6`}/>
                <input
                    id="searchBox"
                    value={searchText}
                    onChange={(e) => 
                        {
                            if(e.target.value.indexOf("\n") == -1) { // check if the user hasn't submitted the search
                                setSearch(e.target.value)
                            } else {
                                // TODO call search function
                            }
                        }}
                    placeholder="Type something..."
                    className={`resize-none whitespace-nowrap overflow-x-scroll [scrollbar-width:none] [-ms-overflow-style:none] [&::-webkit-scrollbar]:hidden overflow-hidden focus:outline-none text-nowrap font-bold font-hack`} 
                />
            </div>
        </div>
    )
}