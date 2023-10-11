### A Pluto.jl notebook ###
# v0.19.29

using Markdown
using InteractiveUtils

# ╔═╡ c2de32e0-6848-11ee-29ed-c59aaa1d23c8
coefficients::Vector{Float64} = [3, 2]

# ╔═╡ 5d135349-5ab6-4447-bf4a-b5b96d76a706
function polyroots(coefficients::Vector{Float64}) :: Vector{Float64}
	count = length(coefficients)
	result::Vector{Float64} = []
	
	if count <= 1
		# not possible
		throw(DomainError("Input vector must contain at least two value."))
	elseif count == 2
		# straight line equation
		if coefficients[2] == 0
			throw(throw(DivideError("polynomial's 1st coeffiecient (e.g. second element in vector) can't be 0")))
		end
		push!(result, -coefficients[1]/coefficients[2])
	elseif count == 3
		# quadratic equation
		if coefficients[2] == 0 and coefficients[3] == 0
			throw(throw(DivideError("polynomial's 1st coeffiecient (e.g. second element in vector) can't be 0")))
			#... okay this aint gonna scale
	end
	return result
end

# ╔═╡ 1a237390-1a8d-4bae-9c89-d207f80a860d
polyroots(coefficients)

# ╔═╡ Cell order:
# ╠═c2de32e0-6848-11ee-29ed-c59aaa1d23c8
# ╠═5d135349-5ab6-4447-bf4a-b5b96d76a706
# ╠═1a237390-1a8d-4bae-9c89-d207f80a860d
