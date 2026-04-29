import { POPPINS } from '@/config/style';

// ── Field ─────────────────────────────────────────────────────────────────────
export function PrimaryField({ label, id, type="text", value, onChange, placeholder, disabled, required=false }) {
	const assocId = id ? id : !label ? '' : label?.toLowerCase().replaceAll(' ', '_');

	return (
		<div>
			{label && (
				<label 
					htmlFor={assocId}
					className="block text-xs font-semibold tracking-widest text-gray-400 uppercase mb-2" 
					style={{ fontFamily: POPPINS}}
				>
					{label}
				</label>
			)}
			<input
				id={assocId}
				type={type}
				value={value}
				onChange={onChange}
				placeholder={placeholder}
				disabled={disabled}
				required={required}
				className="w-full border-0 border-b border-gray-200 pb-2 pt-1 text-sm text-gray-900 bg-transparent outline-none placeholder-gray-300 focus:border-purple-600 transition-colors duration-200 disabled:text-gray-400"
				style={{ fontFamily: POPPINS }}
			/>
		</div>
	);
}