'use client'

import Link from "next/link"

export interface BoardParams {
    name: string
    numNotes: number 
    id: number
}

export default function Board({name, numNotes, id} : BoardParams) {
    return (
        <Link className="aspect-square" href={`/gallery/${id}`}>
            <div className="flex flex-col justify-center items-center aspect-square font-hack font-bold text-text-color ">
                <div className="w-full h-full border-5 rounded-lg flex items-center justify-center">
                    <span>board_image</span>
                </div>
                <div className="flex flex-col items-start w-full">
                    <span className="text-[24px]">{name}</span>
                    <span className="text-[18px]">{numNotes} notes</span>
                </div>
            </div>
        </Link>
    )
}