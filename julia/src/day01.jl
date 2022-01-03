# Based on tutorial at https://viralinstruction.com/posts/aoc2021_1/
module day01

using ..InlineTest

parse(io::IO) = [Base.parse(Int, line, base=10) for line in eachline(io)]
solve(io::IO) = solve(parse(io))
function solve(v::AbstractVector{<:Integer})
    zipdiff(a, b) = sum((@inbounds b[i] > a[i] for i in eachindex(b)), init=0)
    return (zipdiff(v, @view v[2:end]), zipdiff(v, @view v[4:end]))
end

const TEST_STRING = """199
200
208
210
200
207
240
269
260
263"""

@testset "day01" begin
    @test solve(IOBuffer(TEST_STRING)) == (7, 5)
end

end # module


# function parse_input()
#     path = joinpath(@__DIR__, "..", "data", "day01.txt")
#     input = open(path, "r") do file
#         read(file, String)
#     end
#     sonar_reports = parse.(Int, split(input))
#     return sonar_reports
# end

# function part1()
#     ssrs = parse_input()
#     vals = zip(ssrs[1:end-1], ssrs[2:end])
#     return count(v -> (v[1] < v[2]), vals)
# end

# function part2()
#     ssrs = parse_input()
#     triples = collect(zip(ssrs[1:end-2], ssrs[2:end-1], ssrs[3:end]))
#     vals = zip(triples[1:end-1], triples[2:end])
#     return count(v -> (v[1][1] + v[1][2] + v[1][3] < v[2][1] + v[2][2] + v[2][3]), vals)
# end

