def BitmapHoles(strArr):
     # Variables to keep track of visited nodes, and holes 
    visited = set() 
    holes = 0 

    # Iterate through the 2D matrix 
    for i in range(len(strArr)): 
        for j in range(len(strArr[i])): 
  
            # If the current node is 0 and not visited 
            if strArr[i][j] == '0' and (i, j) not in visited: 
  
                # Do a BFS to find the connected component 
                queue = [] 
                queue.append((i, j)) 
  
                # Mark all the nodes in the same component as visited 
                while queue: 
                    x, y = queue.pop(0) 
                    visited.add((x, y)) 
  
                    # Check the top 
                    if x > 0 and strArr[x - 1][y] == '0' and (x - 1, y) not in visited: 
                        queue.append((x - 1, y)) 
  
                    # Check the bottom 
                    if x < len(strArr) - 1 and strArr[x + 1][y] == '0' and (x + 1, y) not in visited: 
                        queue.append((x + 1, y)) 
  
                    # Check the left 
                    if y > 0 and strArr[x][y - 1] == '0' and (x, y - 1) not in visited: 
                        queue.append((x, y - 1)) 
  
                    # Check the right 
                    if y < len(strArr[i]) - 1 and strArr[x][y + 1] == '0' and (x, y + 1) not in visited: 
                        queue.append((x, y + 1)) 
  
                # Increment the number of holes 
                holes += 1 
  
    return holes
