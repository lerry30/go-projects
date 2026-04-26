import { useState, useEffect } from "react";
import { ChevronLeft, ChevronRight } from "lucide-react";

function getPageRange(current, total) {
    if(total <= 7) return Array.from({ length: total }, (_, i) => i + 1);
    if(current <= 3) return [1, 2, 3, 4, "…", total];
    if(current >= total - 2) return [1, "…", total - 3, total - 2, total - 1, total];
    return [1, "…", current - 1, current, current + 1, "…", total];
}

const Pagination = ({ total = 8, defaultPage = 1, callback }) => {
    const [current, setCurrent] = useState(defaultPage);
    const pages = getPageRange(current, total);

    useEffect(() => {
        callback(current);
    }, [current]);

    return (
        <div className="flex items-center gap-1">
            <button
                onClick={() => setCurrent((p) => Math.max(1, p - 1))}
                disabled={current === 1}
                className="flex items-center gap-1.5 px-3 py-1.5 rounded-md border border-gray-200 text-sm font-medium text-gray-500 hover:bg-gray-50 hover:text-gray-800 disabled:opacity-35 disabled:cursor-not-allowed transition-all active:scale-95"
            >
                <ChevronLeft size={14} />
                Previous
            </button>

            {pages.map((p, i) =>
                p === "…" ? (
                    <span key={`ellipsis-${i}`} className="w-8 h-8 flex items-center justify-center text-xs text-gray-400">…</span>
                ) : (
                    <button
                        key={p}
                        onClick={() => setCurrent(p)}
                        className={`w-8 h-8 flex items-center justify-center rounded-md text-sm font-medium transition-all active:scale-95 ${p === current
                                ? "bg-gray-900 text-white"
                                : "text-gray-500 hover:bg-gray-100 hover:text-gray-900"
                            }`}
                    >
                        {p}
                    </button>
                )
            )}

            <button
                onClick={() => setCurrent((p) => Math.min(total, p + 1))}
                disabled={current === total}
                className="flex items-center gap-1.5 px-3 py-1.5 rounded-md border border-gray-200 text-sm font-medium text-gray-500 hover:bg-gray-50 hover:text-gray-800 disabled:opacity-35 disabled:cursor-not-allowed transition-all active:scale-95"
            >
                Next
                <ChevronRight size={14} />
            </button>
        </div>
    );
}

export default Pagination;