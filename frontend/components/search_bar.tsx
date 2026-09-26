'use client'
import SearchIcon from "@/public/search.svg"
import { useState } from "react"

interface SearchBarProps {
    background?: boolean
    full?: boolean
    placeholder?: string
}

export default function SearchBar({ background = false, full = false, placeholder = "Search..." } : SearchBarProps) {
    const [searchText, setSearch] = useState("")
    const backgroundColor = background ? "bg-background-color backdrop-brightness-125" : ""
    const brightness = background ? "backdrop-brightness-125" : ""
    const width = full ? "w-full" : "w-[50%] max-w-[25vw]"
    const justify = full ? "justify-start" : "justify-end"

    // TODO: Create search drop down
    // TODO: REACH GOAL: Make it look like the pintrest dropdown
    return (
        <div className={`${width} h-8 ${backgroundColor} rounded-sm`}>
            <div className={`flex flex-row ${justify} items-center fill-icon-color hover:fill-text-color gap-3 h-8 ${brightness} ${background ? "pl-2" : ""}`}>
                <SearchIcon className={`w-6 h-6`}/>
                <div className={full ? "w-full" : "w-[50%]"}>
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
                    placeholder={placeholder}
                    className={`w-full whitespace-nowrap overflow-x-scroll [scrollbar-width:none] [-ms-overflow-style:none] [&::-webkit-scrollbar]:hidden overflow-hidden focus:outline-none text-nowrap font-bold font-hack`} 
                    />
                </div>
            </div>
        </div>
    )
}