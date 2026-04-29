import { POPPINS, GRADIENT } from '@/config/style';
import { ChevronLeft, ChevronRight } from 'lucide-react';

// ── CTA Button ────────────────────────────────────────────────────────────────
export function PlainButton({ children, onClick, disabled, type = "button", className = "" }) {
	return (
		<button
			type={type}
			onClick={onClick}
			disabled={disabled}
			className={`py-3.5 bg-gray-200 text-gray-700 rounded-2xl text-sm font-semibold tracking-wide cursor-pointer transition-all duration-200 hover:opacity-90 hover:-translate-y-0.5 active:scale-95 disabled:opacity-50 disabled:cursor-not-allowed hover:text-gray-900 ${className}`}
			style={{ boxShadow: "0 8px 24px rgba(60,58,237,0.3)", fontFamily: POPPINS }}
		>
			{children}
		</button>
	);
}

export function GradientButton({ children, onClick, disabled, type = "button", className = "", outlined = false }) {
	if (outlined) {
		return (
			<button
				onClick={onClick}
				className={`relative inline-flex items-center justify-center gap-2 px-6 py-3 rounded-2xl font-semibold text-sm border-2 border-purple-600 text-purple-600 hover:bg-purple-50 transition-all duration-200 ${className}`}
				style={{ fontFamily: POPPINS }}
			>
				{children}
			</button>
		);
	}
	return (
		<button
			onClick={onClick}
			className={`relative inline-flex items-center justify-center gap-2 px-6 py-3.5 rounded-2xl font-semibold text-sm text-white shadow-lg hover:shadow-purple-400/40 hover:-translate-y-0.5 active:scale-95 transition-all duration-200 ${className}`}
			style={{ background: GRADIENT, fontFamily: POPPINS }}
		>
			{children}
		</button>
	);
}