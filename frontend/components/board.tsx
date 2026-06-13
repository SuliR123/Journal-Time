'use client'

export interface BoardParams {
    name: string
    numNotes: number 
    id: number
}

export default function Board({name, numNotes, id} : BoardParams) {
    return (
        <div className="flex flex-col justify-center items-center aspect-square font-hack font-bold">
            <div className="w-full h-full border-5 text-text-color rounded-lg flex items-center justify-center">
                <span>board_image</span>
            </div>
            <div className="flex flex-col items-start w-full">
                <span className="text-[24px]">{name}</span>
                <span className="text-[18px]">{numNotes} notes</span>
            </div>
        </div>
    )
}