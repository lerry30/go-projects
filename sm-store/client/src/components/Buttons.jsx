import { poppins, gradient } from '@/config/style';

// ── CTA Button ────────────────────────────────────────────────────────────────
export function GradientButton({ children, onClick, disabled, type = "button", className = "" }) {
	return (
		<button
			type={type}
			onClick={onClick}
			disabled={disabled}
			className={`py-3.5 rounded-2xl text-white text-sm font-semibold tracking-wide cursor-pointer transition-all duration-200 hover:opacity-90 hover:-translate-y-0.5 active:scale-95 disabled:opacity-50 disabled:cursor-not-allowed ${className}`}
			style={{ background: gradient, boxShadow: "0 8px 24px rgba(124,58,237,0.3)", fontFamily: poppins }}
		>
			{children}
		</button>
	);
}