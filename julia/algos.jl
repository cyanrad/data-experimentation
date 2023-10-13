function euclidGreatestCommonDivider(m::Int64, n::Int64)::Int64
    r = m % n
    return if r == 0
        n
    else
        euclidGreatestCommonDivider(n, r)
    end
end
