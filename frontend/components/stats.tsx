
interface StatsPageProps {
    seconds: number
    wordCount: number
}

function formatZeros(n: number): string {
    if (n < 10 ) {
        return `0${n}`
    } else {
        return `${n}`
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

export default function StatsPage({seconds, wordCount} : StatsPageProps) {
    return(
        <div className="flex flex-col justify-evenly items-start w-full h-full text-[24px] font-hack">
            <span>{`WPM: ${wpm(wordCount, seconds)}`}</span>
            <span>{`Elapsed Time: ${formatTime(seconds)}`}</span>
            <span>{`Word Count: ${wordCount}`}</span>
        </div>
    )
}