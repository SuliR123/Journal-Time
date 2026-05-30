import Scroll from "./scroll";
import SearchBar from "./search_bar";

export default function BoardPopup() {
    return (
        <div className="h-full w-full flex flex-col justify-center items-center">
            <div className="h-full w-full h-full flex flex-col gap-6 items-end rounded-lg">
                <SearchBar />
                <Scroll> {/*TODO: LOAD REAL INFO HERE + CHANGE TO REAL BOARD CARDS */}
                    <div className="flex-shrink-0 w-64 h-32 bg-bg-color text-text-color font-hack font-bold rounded-lg flex items-center justify-center">Item 1</div>
                    <div className="shrink-0 w-64 h-32 bg-blue-300 rounded-lg flex items-center justify-center">Item 2</div>
                    <div className="shrink-0 w-64 h-32 bg-blue-400 rounded-lg flex items-center justify-center">Item 3</div>
                    <div className="shrink-0 w-64 h-32 bg-blue-500 rounded-lg flex items-center justify-center">Item 4</div>
                    <div className="shrink-0 w-64 h-32 bg-blue-600 rounded-lg flex items-center justify-center">Item 5</div>
                </Scroll>
            </div>
        </div>
        
    )
}