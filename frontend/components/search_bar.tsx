import SearchIcon from "@/public/search.svg"
import { useState } from "react"

export default function SearchBar() {
    const [searchText, setSearch] = useState("")

    return (
        <div className="flex flex-row fill-icon-color hover:fill-text-color gap-3">
            <SearchIcon className="w-6 h-6"/>
            <textarea
                id="searchBox"
                value={searchText}
                onChange={(e) => 
                    {
                        if(e.target.value.indexOf("\n") == -1) {
                            setSearch(e.target.value)
                        }
                    }}
                placeholder="Type something..."
                className={`h-6 [scrollbar-width:none] [-ms-overflow-style:none] [&::-webkit-scrollbar]:hidden overflow-hidden focus:outline-none text-nowrap font-bold font-hack`} 
            />
        </div>
    )
}