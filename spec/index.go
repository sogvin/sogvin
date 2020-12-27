package spec

import (
	. "github.com/gregoryv/web"
)

func NewExploreRequirementsEngineering(n *Hn) *Element {
	return Article(
		n.H1("Exploring requirements engineering"),

		P(`An exercise in elicitating requirements, imho. still one of
        the most difficult task in software engineering.`),

		P(`As a software engineer you are tasked to produce software
		systems to fulfill the need of a stakeholder. I use the term
		software engineer, or just engineer, for all roles used today
		in the industry that somehow contribute to producing
		software. The reason is they all have one thing incommon, they
		have to understand the purpose of their work. Without it, the
		end result will never be as good as envisioned by the
		stakeholder.`),

		P(`As an engineer I solve problems. One reoccuring problem is
		the difficulty of conveying knowledge from stakeholders to the
		engineer, in such a manner that it is easily understood.
		There are many reasons for this and hopefully with this
		exercise I'll highlight some of them and provide some
		solutions, being an engineer and all.`),

		P(`Throughout this exercise you will follow along a fiction
		story of an enterprise developing a space ship control
		system. In parts I'll use dialog form between stakeholders and
		engineers to highlight the iterative process required to
		produce easily digested specifications and requirements for
		software developers in particular. The specification can and
		is often a base which agreements are founded upon, so all
		stakeholders should be able to digest it easily, not only
		developers.`),
		//
	)
}
