import Board from "@/components/board";
import {BoardParams} from "@/components/board"
import Scroll from "@/components/scroll";
import SearchBar from "@/components/search_bar";

export default function Gallery() {

    const dummyData = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16]

    return (
        <div className="flex flex-col w-full h-[100vh] items-center">
            <div className="flex flex-col items-start w-[85%]">
                <div className="h-[4vh]"/>
                <SearchBar full={true} background={true} placeholder="Search your boards"/>
            </div>
            <div className="flex flex-col font-hack font-bold w-full h-full justify-center items-center">
                <div className="flex flex-col w-[85%] h-[90%] justify-between items-start text-[32px] gap-4">
                    <div className="flex flex-row justify-between items-center w-full">
                        <span>Name's Gallery</span>
                        <span>Create +</span>
                    </div>
                    <div className="h-[90%]">
                        <Scroll vertical={true}> {/*TODO: LOAD REAL INFO HERE + CHANGE TO REAL BOARD CARDS */}
                            {dummyData.map((value: number, index: number) => 
                                <Board name={`Board ${value}`} numNotes={10} id={index} key={index}/>
                            )} {/* Display number of notes in the board, other metadata like date created maybe*/}
                        </Scroll>
                    </div>
                    
                </div>
            </div>
        </div>
    )
}