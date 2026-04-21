import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { SuccessModal, ErrorModal } from '@/components/Modal';
import { sendJSON } from '@/utils/send';
import { BASE_URL } from '@/config/server';

import { GradientButton } from '@/components/Buttons';
import { PrimaryField } from '@/components/Fields';

import IllustrationSVG from '@/components/IllustrationSVG';
import Loading from '@/components/Loading'

const StarIcon = () => (
	<svg width="14" height="14" viewBox="0 0 24 24" fill="white">
		<path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z" />
	</svg>
);

const Signin = () => {
	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");
	const [remember, setRemember] = useState(false);
	const [loading, setLoading] = useState(false);
	const [modal, setModal] = useState({message: '', type: null});

	const navigate = useNavigate();

	const handleSubmit = async (e) => {
		try {
			e.preventDefault();
			setLoading(true);

			const tForm = {
				email: String(email).trim(),
				password: String(password).trim()
			};

			if(!tForm.email || !tForm.password)
				throw new Error('Fields are required');
			
			setLoading(true);
			const result = await sendJSON(`${BASE_URL}/signin`, tForm);
			if(result) {
				setModal({message: 'Login successfully.', type: 'success'})
				setTimeout(() => {
					setModal({message: '', type: null});
					navigate('/');
				}, 1000);
			}
		} catch(err) {
			const errorMessage = err?.message || 'Something went wrong';
			console.log("Failed to signin: ", errorMessage);
			setModal({message: errorMessage, type: 'error'});
			setTimeout(() => {
				setModal({message: '', type: null});
			}, 1000);
		} finally {
			setLoading(false);
		}
	};

	return (
		<div
			className="min-h-screen flex items-center justify-center bg-gray-100 p-4"
			style={{ fontFamily: "'Poppins', sans-serif" }}
		>
			{/* Import Poppins */}
			<style>{`@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;500;600;700&display=swap');`}</style>

			<div className="flex w-full max-w-5xl rounded-3xl overflow-hidden shadow-2xl bg-white"
				style={{ minHeight: "620px" }}>

				{/* ── LEFT PANEL ── */}
				<div className="flex flex-col justify-center w-full lg:w-1/2 px-10 py-14 bg-white">

					{/* Logo */}
					<div className="flex items-center gap-2 mb-12">
						<div
							className="w-8 h-8 rounded-xl flex items-center justify-center"
							style={{ background: "linear-gradient(135deg, #7C3AED, #C026D3)" }}
						>
							<StarIcon />
						</div>
						<span className="text-base font-semibold text-gray-900 tracking-tight">Luminary</span>
					</div>

					{/* Heading */}
					<h1 className="text-3xl font-bold text-gray-900 leading-tight tracking-tight mb-1">
						Welcome back
					</h1>
					<p className="text-sm text-gray-400 mb-9">Sign in to your account to continue</p>

					{/* Form */}
					<form onSubmit={handleSubmit} className="space-y-6">
						{/* Email */}
						<PrimaryField 
							label="Email Address"
							type="email"
							value={email}
							onChange={(e) => setEmail(e.target.value)}
							placeholder="you@email.com"
							required
						/>

						{/* Password */}
						<PrimaryField
							label="Password"
							type="password"
							value={password}
							onChange={(e) => setPassword(e.target.value)}
							placeholder="••••••••••"
							required
						/>

						{/* Remember + Forgot */}
						<div className="flex items-center justify-between">
							<label className="flex items-center gap-2 cursor-pointer text-xs text-gray-400">
								<input
									type="checkbox"
									checked={remember}
									onChange={(e) => setRemember(e.target.checked)}
									className="accent-purple-600 w-3.5 h-3.5"
								/>
								Remember me
							</label>
							<a href="#" className="text-xs font-semibold text-purple-600 hover:text-purple-700 transition-colors">
								Forgot password?
							</a>
						</div>

						{/* CTA Button */}
						<GradientButton
							type="submit"
							disabled={loading}
							className="w-full"
						>
							{loading ? "Signing in..." : "Sign in to dashboard →"}
						</GradientButton>
					</form>

					{/* Sign up link */}
					<p className="text-center text-xs text-gray-400 mt-6">
						Don't have an account?{" "}
						<Link to="/signup" className="text-purple-600 font-semibold hover:text-purple-700 transition-colors">
							Create one free →
						</Link>
					</p>
				</div>

				{/* ── RIGHT PANEL ── */}
				<div
					className="hidden lg:flex lg:w-1/2 relative overflow-hidden"
					style={{
						background: "linear-gradient(160deg, #4f1899 0%, #7C3AED 30%, #C026D3 62%, #f97316 100%)",
					}}
				>
					{/* Blob decorations */}
					<div
						className="absolute -top-16 -right-16 w-80 h-80 rounded-full"
						style={{ background: "rgba(255,255,255,0.07)" }}
					/>
					<div
						className="absolute -bottom-20 -left-20 w-96 h-96 rounded-full"
						style={{ background: "rgba(255,150,50,0.13)" }}
					/>
					<div
						className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-64 h-64 rounded-full"
						style={{ background: "rgba(255,255,255,0.05)" }}
					/>

					{/* Content */}
					<div className="relative z-10 flex flex-col justify-center px-12 py-14">

						{/* Sale badge */}
						<div
							className="inline-flex items-center gap-2 self-start rounded-full px-4 py-1.5 mb-8 text-xs font-medium text-white/90 border border-white/20"
							style={{ background: "rgba(255,255,255,0.12)", backdropFilter: "blur(4px)" }}
						>
							<span className="w-1.5 h-1.5 rounded-full bg-green-400 shadow-sm shadow-green-300" />
							Free shipping on orders over ₱999
						</div>

						<h2 className="text-2xl font-bold text-white leading-snug tracking-tight mb-3">
							Everything you need,<br />delivered to your door.
						</h2>
						<p className="text-sm text-white/70 leading-relaxed mb-8">
							Shop from 50,000+ products across fashion, tech, and home essentials.
							Same-day delivery available in Metro Manila.
						</p>

						<IllustrationSVG />

						{/* Stat cards */}
						<div className="flex gap-3 mt-7">
							{[
								{ num: "50K+", label: "Products listed" },
								{ num: "1-Day", label: "Avg. delivery" },
								{ num: "4.8★", label: "Seller rating" },
							].map(({ num, label }) => (
								<div
									key={label}
									className="flex-1 rounded-2xl px-4 py-3 border border-white/20"
									style={{ background: "rgba(255,255,255,0.1)", backdropFilter: "blur(6px)" }}
								>
									<div className="text-xl font-bold text-white tracking-tight">{num}</div>
									<div className="text-xs text-white/60 mt-0.5">{label}</div>
								</div>
							))}
						</div>
					</div>
				</div>
			</div>

			{loading && <Loading />}

			{(modal?.message && modal?.type === 'success') && (
				<SuccessModal
					message={modal.message}
					callback={() => {
						setModal({message: '', type: null});
						navigate('/')
					}}
				/>
			)}

			{(modal?.message && modal?.type === 'error') && (
				<ErrorModal
					header="Oops!"
					message={modal.message}
					callback={() => {
						setModal({message: '', type: null});
					}}
				/>
			)}
		</div>
	);
}

export default Signin;