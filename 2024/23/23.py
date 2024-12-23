import networkx as nx
G = nx.Graph()
# G.add_edges_from
G.add_edges_from([("kh","tc"),("qp","kh"),("de","cg"),("ka","co"),("yn","aq"),("qp","ub"),("cg","tb"),("vc","aq"),("tb","ka"),("wh","tc"),("yn","cg"),("kh","ub"),("ta","co"),("de","co"),("tc","td"),("tb","wq"),("wh","td"),("ta","ka"),("td","qp"),("aq","cg"),("wq","ub"),("ub","vc"),("de","ta"),("wq","aq"),("wq","vc"),("wh","yn"),("ka","de"),("kh","ta"),("co","tc"),("wh","qp"),("tb","vc"),("td","yn")])
cliques = list(nx.find_cliques(G))
print(sorted(max(cliques, key=len)))
